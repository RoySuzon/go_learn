package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ---------------------------------------------------------
// 1. DATA MODELS (User, Book, Claims)
// ---------------------------------------------------------

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Stored as Hashed Password
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"` // Expiration Timestamp
}

// ---------------------------------------------------------
// 2. IN-MEMORY STORAGE & SECRET KEY
// ---------------------------------------------------------

const jwtSecretKey = "super-secret-key-12345"

var (
	users = []User{
		{ID: 1, Username: "goutom", Password: hashPassword("secret123")},
	}
	books = []Book{
		{ID: 1, Title: "Go REST API & JWT Masterclass", Author: "Goutom Roy"},
		{ID: 2, Title: "Concurrent Systems in Go", Author: "Sujan Roy"},
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

// Generates Base64URL-encoded JWT (HS256)
func generateJWT(userID int, username string) (string, error) {
	// Header
	headerJSON, _ := json.Marshal(map[string]string{"alg": "HS256", "typ": "JWT"})
	headerB64 := base64.RawURLEncoding.EncodeToString(headerJSON)

	// Payload (1 Hour expiration)
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Exp:      time.Now().Add(1 * time.Hour).Unix(),
	}
	payloadJSON, _ := json.Marshal(claims)
	payloadB64 := base64.RawURLEncoding.EncodeToString(payloadJSON)

	// Signature
	unsignedToken := headerB64 + "." + payloadB64
	mac := hmac.New(sha256.New, []byte(jwtSecretKey))
	mac.Write([]byte(unsignedToken))
	signatureB64 := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return unsignedToken + "." + signatureB64, nil
}

// Validates JWT Signature and Expiration
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
// 4. AUTHENTICATION HANDLERS (Register, Login, Profile)
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

	dbMutex.Lock()
	defer dbMutex.Unlock()

	for _, u := range users {
		if u.Username == req.Username {
			http.Error(w, `{"error": "Username already exists"}`, http.StatusConflict)
			return
		}
	}

	newUser := User{
		ID:       len(users) + 1,
		Username: req.Username,
		Password: hashPassword(req.Password),
	}
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully!",
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

	hashedInputPassword := hashPassword(req.Password)

	dbMutex.Lock()
	var foundUser *User
	for _, u := range users {
		if u.Username == req.Username && u.Password == hashedInputPassword {
			foundUser = &u
			break
		}
	}
	dbMutex.Unlock()

	if foundUser == nil {
		http.Error(w, `{"error": "Invalid username or password"}`, http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(foundUser.ID, foundUser.Username)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful!",
		"token":   token,
	})
}

// ---------------------------------------------------------
// 5. PROTECTED MIDDLEWARE & HANDLERS
// ---------------------------------------------------------

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

		// Proceed to handler
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

func booksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dbMutex.Lock()
	defer dbMutex.Unlock()
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/api/register", registerHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/books", booksHandler)
	http.HandleFunc("/api/profile", authMiddleware(profileHandler))

	fmt.Println("==================================================")
	fmt.Println(" 🚀 REST API with JWT Auth Server is Running!")
	fmt.Println(" 📍 Public Endpoints:")
	fmt.Println("    - POST /api/register")
	fmt.Println("    - POST /api/login")
	fmt.Println("    - GET  /api/books")
	fmt.Println(" 🔒 Protected Endpoints (Requires Bearer Token):")
	fmt.Println("    - GET  /api/profile")
	fmt.Println("==================================================")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
