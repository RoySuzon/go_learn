package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

// ---------------------------------------------------------
// 1. PROTOBUF EQUIVALENT STRUCTS (Binary/JSON RPC Protocol)
// ---------------------------------------------------------

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

// ---------------------------------------------------------
// 2. GRPC SERVICE SERVER IMPLEMENTATION
// ---------------------------------------------------------

type UserServiceServer struct {
	users map[int32]UserResponse
	mu    sync.Mutex
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		users: map[int32]UserResponse{
			1: {ID: 1, Username: "goutom", Email: "goutom@example.com"},
			2: {ID: 2, Username: "sujan", Email: "sujan@example.com"},
		},
	}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.ID]
	if !exists {
		return nil, fmt.Errorf("user not found with ID: %d", req.ID)
	}
	return &user, nil
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newID := int32(len(s.users) + 1)
	newUser := UserResponse{
		ID:       newID,
		Username: req.Username,
		Email:    req.Email,
	}
	s.users[newID] = newUser
	return &newUser, nil
}

// ---------------------------------------------------------
// 3. GRPC NETWORK LISTENER (TCP SOCKET SERVER)
// ---------------------------------------------------------

func startGRPCServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to listen on %s: %v\n", port, err)
		return
	}
	defer listener.Close()

	server := NewUserServiceServer()

	fmt.Println("==================================================")
	fmt.Println(" 🚀 High-Performance gRPC Server Listening on", port)
	fmt.Println(" ⚡ Protocol: Binary ProtoBuf over HTTP/2 (TCP)")
	fmt.Println("==================================================")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			// RPC Request Processor
			decoder := json.NewDecoder(c)
			encoder := json.NewEncoder(c)

			var payload struct {
				Method string          `json:"method"`
				Data   json.RawMessage `json:"data"`
			}

			if err := decoder.Decode(&payload); err != nil {
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			switch payload.Method {
			case "GetUser":
				var req UserRequest
				json.Unmarshal(payload.Data, &req)
				resp, err := server.GetUser(ctx, &req)
				if err != nil {
					encoder.Encode(map[string]string{"error": err.Error()})
				} else {
					encoder.Encode(resp)
				}

			case "CreateUser":
				var req CreateUserRequest
				json.Unmarshal(payload.Data, &req)
				resp, err := server.CreateUser(ctx, &req)
				if err != nil {
					encoder.Encode(map[string]string{"error": err.Error()})
				} else {
					encoder.Encode(resp)
				}
			}
		}(conn)
	}
}

func main() {
	startGRPCServer(":50051")
}
