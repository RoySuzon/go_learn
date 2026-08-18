package main

import "fmt"

// Package-level constants using iota generator
const (
	StatusPending  = iota // 0
	StatusApproved        // 1
	StatusRejected        // 2
)

// Package-level constant
const AppVersion = "v1.0.0"

func main() {
	fmt.Println("=== 2.1 Variable Declarations ===")
	// 1. Explicit variable declaration
	var username string = "Goutom"

	// 2. Type inference (Go infers type based on value)
	var age = 26

	// 3. Short variable declaration syntax := (Functions only)
	score := 98.75

	// 4. Multiple variables in a single line
	var x, y int = 10, 20
	city, country := "Dhaka", "Bangladesh"

	fmt.Printf("User: %s (Age: %d) | Score: %.2f\n", username, age, score)
	fmt.Printf("Location: %s, %s (Coordinates: %d, %d)\n", city, country, x, y)

	fmt.Println("\n=== 2.2 Zero Values ===")
	// Variables declared without values get default "zero values"
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Printf("int default: %v\n", defaultInt)         // 0
	fmt.Printf("float default: %v\n", defaultFloat)     // 0
	fmt.Printf("bool default: %v\n", defaultBool)       // false
	fmt.Printf("string default: %q\n", defaultString)   // ""

	fmt.Println("\n=== 2.3 Constants & iota ===")
	fmt.Println("App Version:", AppVersion)
	fmt.Printf("Statuses -> Pending: %d, Approved: %d, Rejected: %d\n",
		StatusPending, StatusApproved, StatusRejected)

	fmt.Println("\n=== 2.4 Explicit Type Conversion ===")
	var intVal int = 42
	var floatVal float64 = float64(intVal) // Explicit conversion required
	var uintVal uint = uint(floatVal)

	fmt.Printf("Original int: %d (%T)\n", intVal, intVal)
	fmt.Printf("Converted float64: %.2f (%T)\n", floatVal, floatVal)
	fmt.Printf("Converted uint: %d (%T)\n", uintVal, uintVal)
}
