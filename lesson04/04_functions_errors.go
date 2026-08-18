package main

import (
	"errors"
	"fmt"
)

// 1. Basic Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// 2. Named Return Values & Naked Return
func getRectangleStats(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Returns 'area' and 'perimeter' automatically
}

// 3. Variadic Function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 4. Closure Generator
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 5. Function demonstrating Defer, Panic, and Recover
func safeExecution() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("--> Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting safeExecution routine...")
	panic("unexpected system outage!") // Triggers panic
	// fmt.Println("Unreachable code")
}

func main() {
	fmt.Println("=== 4.1 Multiple Return Values & Errors ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	_, errZero := divide(10, 0)
	if errZero != nil {
		fmt.Println("Handled Error:", errZero)
	}

	fmt.Println("\n=== 4.2 Named Returns & Variadic Functions ===")
	a, p := getRectangleStats(5.0, 3.0)
	fmt.Printf("Rectangle Stats -> Area: %.1f, Perimeter: %.1f\n", a, p)

	totalSum := sum(10, 20, 30, 40)
	fmt.Println("Variadic Sum (10+20+30+40):", totalSum)

	fmt.Println("\n=== 4.3 Closures & Anonymous Functions ===")
	counterA := makeCounter()
	fmt.Println("Counter A Call 1:", counterA()) // 1
	fmt.Println("Counter A Call 2:", counterA()) // 2

	counterB := makeCounter()
	fmt.Println("Counter B Call 1:", counterB()) // 1 (independent state)

	fmt.Println("\n=== 4.4 Defer, Panic & Recover ===")
	// Defer statements execute in LIFO (Last-In-First-Out) order when main returns
	defer fmt.Println("--> Deferred cleanup task (Executed when main completes)")

	safeExecution()
}
