package main

import (
	"context"
	"fmt"
	"time"
)

// Simulated slow database call honoring context deadline
func fetchUserData(ctx context.Context, userID int) (string, error) {
	select {
	case <-time.After(500 * time.Millisecond): // Takes 500ms
		return fmt.Sprintf("User details for ID: %d", userID), nil
	case <-ctx.Done(): // Context Timeout or Cancellation
		return "", ctx.Err()
	}
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 Lesson 12: Context Package (Timeout & Cancellation)")
	fmt.Println("==================================================")

	// 1. Context with Value (Trace ID / Request ID)
	type key string
	const requestIDKey key = "request_id"
	ctxWithValue := context.WithValue(context.Background(), requestIDKey, "REQ-99812")
	fmt.Println("১. Request ID Context Value:", ctxWithValue.Value(requestIDKey))

	// 2. Context with Timeout (200ms Timeout vs 500ms DB task -> Expect Timeout)
	fmt.Println("\n২. Context Timeout পরীক্ষা (Timeout 200ms, DB Task 500ms):")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelTimeout()

	data, err := fetchUserData(ctxTimeout, 101)
	if err != nil {
		fmt.Println("   ❌ Call Cancelled by Context:", err) // context deadline exceeded
	} else {
		fmt.Println("   ✅ Fetched Data:", data)
	}

	// 3. Context with Sufficient Time (800ms Timeout vs 500ms DB task -> Expect Success)
	fmt.Println("\n৩. পর্যাপ্ত সময় সহ Context পরীক্ষা (Timeout 800ms):")
	ctxSuccess, cancelSuccess := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancelSuccess()

	dataSuccess, errSuccess := fetchUserData(ctxSuccess, 102)
	if errSuccess != nil {
		fmt.Println("   ❌ Error:", errSuccess)
	} else {
		fmt.Println("   ✅ Data Received Successfully:", dataSuccess)
	}
}
