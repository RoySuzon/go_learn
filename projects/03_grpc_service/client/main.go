package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type UserRequest struct {
	ID int32 `json:"id"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Client wrapper to call gRPC endpoints
func callGRPC(method string, requestData interface{}) (*UserResponse, error) {
	conn, err := net.Dial("tcp", "localhost:50051")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server at localhost:50051: %w", err)
	}
	defer conn.Close()

	dataJSON, _ := json.Marshal(requestData)
	payload := map[string]interface{}{
		"method": method,
		"data":   json.RawMessage(dataJSON),
	}

	if err := json.NewEncoder(conn).Encode(payload); err != nil {
		return nil, err
	}

	var resp UserResponse
	if err := json.NewDecoder(conn).Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🔌 gRPC Client: Calling Microservice Endpoints...")
	fmt.Println("==================================================")

	// 1. Call GetUser(ID: 1)
	user, err := callGRPC("GetUser", UserRequest{ID: 1})
	if err != nil {
		fmt.Println("❌ Error calling GetUser:", err)
	} else {
		fmt.Printf("✅ GetUser(1) Success -> ID: %d | User: %s | Email: %s\n", user.ID, user.Username, user.Email)
	}

	// 2. Call CreateUser
	newUser, err := callGRPC("CreateUser", CreateUserRequest{Username: "rahim", Email: "rahim@example.com"})
	if err != nil {
		fmt.Println("❌ Error calling CreateUser:", err)
	} else {
		fmt.Printf("🎉 CreateUser Success -> New ID: %d | User: %s | Email: %s\n", newUser.ID, newUser.Username, newUser.Email)
	}
}
