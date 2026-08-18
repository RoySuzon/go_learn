package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 3.1 Conditional Statements (if/else) ===")
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A+")
	} else if score >= 80 {
		fmt.Println("Grade: A")
	} else {
		fmt.Println("Grade: B or below")
	}

	// If statement with short variable initialization
	// 'maxLimit' is scoped only within this if-else block
	if maxLimit := 100; score < maxLimit {
		fmt.Printf("Score (%d) is valid within max limit (%d)\n", score, maxLimit)
	}

	fmt.Println("\n=== 3.2 Switch Statements ===")
	day := "Wednesday"

	// 1. Expression Switch (Notice: No 'break' needed in Go!)
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println(day, "is a Weekday.")
	case "Saturday", "Sunday":
		fmt.Println(day, "is a Weekend!")
	default:
		fmt.Println("Invalid day.")
	}

	// 2. Tagless Switch (Acts like an if-else chain)
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Greeting: Good Morning!")
	case hour < 17:
		fmt.Println("Greeting: Good Afternoon!")
	default:
		fmt.Println("Greeting: Good Evening!")
	}

	fmt.Println("\n=== 3.3 For Loops ===")
	// Go has ONLY one loop keyword: 'for'

	// 1. Standard 3-component loop
	fmt.Print("Standard loop (1 to 3): ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 2. While-style loop (condition only)
	fmt.Print("While-style loop (3 down to 1): ")
	counter := 3
	for counter > 0 {
		fmt.Printf("%d ", counter)
		counter--
	}
	fmt.Println()

	// 3. Infinite loop with break and continue
	fmt.Print("Infinite loop with break/continue: ")
	n := 0
	for {
		n++
		if n%2 == 0 {
			continue // Skip even numbers
		}
		if n > 5 {
			break // Exit loop when n exceeds 5
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
