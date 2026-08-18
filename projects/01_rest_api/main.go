package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

// ---------------------------------------------------------
// 1. GORM DATA MODELS FOR POSTGRESQL
// ---------------------------------------------------------

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Author    string    `gorm:"not null" json:"author"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

// ---------------------------------------------------------
// 2. IN-MEMORY FALLBACK DB & SECRET KEY
// ---------------------------------------------------------

const jwtSecretKey = "super-secret-key-12345"

var (
	fallbackUsers = []User{
		{ID: 1, Username: "goutom", Password: hashPassword("secret123")},
	}
	fallbackBooks = []Book{
		{ID: 1, Title: "Go REST API & GORM PostgreSQL", Author: "Goutom Roy", Price: 450},
		{ID: 2, Title: "Concurrent Systems in Go", Author: "Sujan Roy", Price: 380},
	}
	dbMutex sync.Mutex
)

// ---------------------------------------------------------
// 3. CRYPTO & JWT HELPER FUNCTIONS
// ---------------------------------------------------------

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password + "salt_key_99"))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func generateJWT(userID uint, username string) (string, error) {
	headerJSON, _ := json.Marshal(map[string]string{"alg": "HS256", "typ": "JWT"})
	headerB64 := base64.RawURLEncoding.EncodeToString(headerJSON)

	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Exp:      time.Now().Add(1 * time.Hour).Unix(),
	}
	payloadJSON, _ := json.Marshal(claims)
	payloadB64 := base64.RawURLEncoding.EncodeToString(payloadJSON)

	unsignedToken := headerB64 + "." + payloadB64
	mac := hmac.New(sha256.New, []byte(jwtSecretKey))
	mac.Write([]byte(unsignedToken))
	signatureB64 := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return unsignedToken + "." + signatureB64, nil
}

func validateJWT(tokenString string) (*JWTClaims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	unsignedToken := parts[0] + "." + parts[1]
	mac := hmac.New(sha256.New, []byte(jwtSecretKey))
	mac.Write([]byte(unsignedToken))
	expectedSignature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
		return nil, fmt.Errorf("invalid signature")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid payload encoding")
	}

	var claims JWTClaims
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, fmt.Errorf("invalid claims")
	}

	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("token has expired")
	}

	return &claims, nil
}

// ---------------------------------------------------------
// 4. API HANDLERS (POSTGRESQL + GORM)
// ---------------------------------------------------------

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Username == "" || req.Password == "" {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	hashedPass := hashPassword(req.Password)

	// If PostgreSQL GORM is connected
	if DB != nil {
		var count int64
		DB.Model(&User{}).Where("username = ?", req.Username).Count(&count)
		if count > 0 {
			http.Error(w, `{"error": "Username already exists in PostgreSQL"}`, http.StatusConflict)
			return
		}

		user := User{Username: req.Username, Password: hashedPass}
		if err := DB.Create(&user).Error; err != nil {
			http.Error(w, `{"error": "Failed to save user to PostgreSQL"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User registered in PostgreSQL DB!",
			"user_id": user.ID,
		})
		return
	}

	// In-memory fallback
	dbMutex.Lock()
	defer dbMutex.Unlock()
	for _, u := range fallbackUsers {
		if u.Username == req.Username {
			http.Error(w, `{"error": "Username already exists"}`, http.StatusConflict)
			return
		}
	}

	newUser := User{ID: uint(len(fallbackUsers) + 1), Username: req.Username, Password: hashedPass}
	fallbackUsers = append(fallbackUsers, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully (In-Memory)!",
		"user_id": newUser.ID,
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid payload"}`, http.StatusBadRequest)
		return
	}

	hashedPass := hashPassword(req.Password)
	var userID uint
	var found bool

	if DB != nil {
		var user User
		if err := DB.Where("username = ? AND password = ?", req.Username, hashedPass).First(&user).Error; err == nil {
			userID = user.ID
			found = true
		}
	} else {
		dbMutex.Lock()
		for _, u := range fallbackUsers {
			if u.Username == req.Username && u.Password == hashedPass {
				userID = u.ID
				found = true
				break
			}
		}
		dbMutex.Unlock()
	}

	if !found {
		http.Error(w, `{"error": "Invalid username or password"}`, http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(userID, req.Username)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful!",
		"token":   token,
	})
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if DB != nil {
		var bList []Book
		DB.Find(&bList)
		json.NewEncoder(w).Encode(bList)
		return
	}

	dbMutex.Lock()
	defer dbMutex.Unlock()
	json.NewEncoder(w).Encode(fallbackBooks)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, `{"error": "Missing or invalid Authorization header"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := validateJWT(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-User-ID", fmt.Sprintf("%d", claims.UserID))
		r.Header.Set("X-Username", claims.Username)
		next(w, r)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-Username")
	userID := r.Header.Get("X-User-ID")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Welcome to your protected profile!",
		"user_id":  userID,
		"username": username,
	})
}

// ---------------------------------------------------------
// 5. SERVER ENTRYPOINT
// ---------------------------------------------------------

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn != "" {
		_, err := initDatabase(dsn)
		if err != nil {
			fmt.Println("⚠️ PostgreSQL Warning:", err)
			fmt.Println("ℹ️ Falling back to In-Memory storage mode.")
		}
	} else {
		fmt.Println("ℹ️ No POSTGRES_DSN provided. Running in Standalone / In-Memory mode.")
	}

	http.HandleFunc("/api/register", registerHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/books", booksHandler)
	http.HandleFunc("/api/profile", authMiddleware(profileHandler))

	fmt.Println("==================================================")
	fmt.Println(" 🚀 REST API with PostgreSQL & GORM Server Active")
	fmt.Println(" 📍 Endpoints:")
	fmt.Println("    - POST /api/register")
	fmt.Println("    - POST /api/login")
	fmt.Println("    - GET  /api/books")
	fmt.Println("    - GET  /api/profile (Protected)")
	fmt.Println("==================================================")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
