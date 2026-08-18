package main

import (
	"fmt"
	"sync"
	"time"
)

// ---------------------------------------------------------
// 1. TOKEN BUCKET RATE LIMITER PATTERN
// ---------------------------------------------------------

type TokenBucketRateLimiter struct {
	capacity     int
	tokens       int
	refillRate   time.Duration
	lastRefilled time.Time
	mu           sync.Mutex
}

func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefilled: time.Now(),
	}
}

func (tb *TokenBucketRateLimiter) AllowRequest() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	// Refill tokens based on elapsed time
	elapsed := now.Sub(tb.lastRefilled)
	if elapsed >= tb.refillRate {
		tb.tokens = tb.capacity
		tb.lastRefilled = now
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true // Request Allowed
	}
	return false // Rate Limit Exceeded (HTTP 429)
}

// ---------------------------------------------------------
// 2. CIRCUIT BREAKER PATTERN (Closed -> Open -> Half-Open)
// ---------------------------------------------------------

type CircuitBreakerState string

const (
	StateClosed   CircuitBreakerState = "CLOSED"
	StateOpen     CircuitBreakerState = "OPEN"
	StateHalfOpen CircuitBreakerState = "HALF-OPEN"
)

type CircuitBreaker struct {
	failureThreshold int
	failureCount     int
	state            CircuitBreakerState
	lastStateChange  time.Time
	mu               sync.Mutex
}

func NewCircuitBreaker(threshold int) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		state:            StateClosed,
	}
}

func (cb *CircuitBreaker) Execute(reqFunc func() error) error {
	cb.mu.Lock()
	if cb.state == StateOpen {
		if time.Since(cb.lastStateChange) > 2*time.Second {
			cb.state = StateHalfOpen
			fmt.Println("⚡ [CIRCUIT BREAKER] State changed to HALF-OPEN (Testing downstream service)...")
		} else {
			cb.mu.Unlock()
			return fmt.Errorf("circuit breaker is OPEN: Request blocked to protect system")
		}
	}
	cb.mu.Unlock()

	err := reqFunc()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failureCount++
		if cb.failureCount >= cb.failureThreshold {
			cb.state = StateOpen
			cb.lastStateChange = time.Now()
			fmt.Println("🚨 [CIRCUIT BREAKER] State changed to OPEN! Tripping circuit due to high errors.")
		}
		return err
	}

	// Success resets state
	cb.failureCount = 0
	cb.state = StateClosed
	return nil
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 System Design Patterns: Rate Limiter & Circuit Breaker")
	fmt.Println("==================================================")

	// 1. Rate Limiter Demo (Max 3 requests per refill)
	fmt.Println("\n--- 🛡️ Token Bucket Rate Limiter Demo ---")
	limiter := NewTokenBucket(3, 1*time.Second)

	for i := 1; i <= 5; i++ {
		allowed := limiter.AllowRequest()
		if allowed {
			fmt.Printf("Request #%d: ✅ ALLOWED (HTTP 200)\n", i)
		} else {
			fmt.Printf("Request #%d: ❌ BLOCKED (HTTP 429 Too Many Requests)\n", i)
		}
	}

	// 2. Circuit Breaker Demo
	fmt.Println("\n--- ⚡ Circuit Breaker Pattern Demo ---")
	cb := NewCircuitBreaker(2)

	failingServiceCall := func() error {
		return fmt.Errorf("downstream database timeout error")
	}

	for i := 1; i <= 4; i++ {
		err := cb.Execute(failingServiceCall)
		fmt.Printf("Service Call #%d Result -> Error: %v\n", i, err)
	}
}
