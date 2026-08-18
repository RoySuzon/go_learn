package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Main function is the entry point of every executable Go program.
func main() {
	// 1. Basic Output
	fmt.Println("=== 1.1 Hello World & Basic Printing ===")
	fmt.Println("Welcome to Go (Golang) Programming!")

	// 2. Formatted Printing (fmt.Printf)
	// %s = string, %d = integer, %s = string
	fmt.Printf("Go version: %s | Platform: %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	// 3. Packages & Standard Library Functions
	fmt.Println("\n=== 1.2 Using Standard Library Packages ===")
	now := time.Now()
	fmt.Println("Current Time:", now.Format("2006-01-02 15:04:05"))

	radius := 5.0
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Circle radius: %.1f -> Area: %.2f\n", radius, area)
}
