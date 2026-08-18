package main

import (
	"testing"
)

func TestJWTGenerationAndValidation(t *testing.T) {
	var userID uint = 101
	username := "goutom"

	// 1. Generate JWT
	token, err := generateJWT(userID, username)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	if token == "" {
		t.Fatalf("Expected non-empty JWT token string")
	}

	// 2. Validate valid token
	claims, err := validateJWT(token)
	if err != nil {
		t.Fatalf("Failed to validate JWT token: %v", err)
	}

	if claims.UserID != userID || claims.Username != username {
		t.Errorf("Claims mismatch! Expected user %d (%s), got %d (%s)", userID, username, claims.UserID, claims.Username)
	}

	// 3. Test tampered token
	tamperedToken := token + "invalid"
	_, errTampered := validateJWT(tamperedToken)
	if errTampered == nil {
		t.Errorf("Expected error for tampered token, but got nil")
	}
}

func TestPasswordHashing(t *testing.T) {
	pass := "mySecretPass123"
	hash1 := hashPassword(pass)
	hash2 := hashPassword(pass)

	if hash1 != hash2 {
		t.Errorf("Hashes for same password should match")
	}

	if hash1 == pass {
		t.Errorf("Hashed password should not equal raw password")
	}
}
