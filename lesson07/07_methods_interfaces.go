package main

import (
	"fmt"
	"math"
)

// Struct definition
type Circle struct {
	Radius float64
}

// 1. Method with VALUE Receiver (Read-only operation)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 2. Method with POINTER Receiver (Modifies receiver struct)
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

// 3. Interface Definition
// In Go, interfaces are implemented IMPLICITLY!
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Helper function accepting ANY type satisfying Shape interface
func printArea(s Shape) {
	fmt.Printf("Area of %T: %.2f\n", s, s.Area())
}

// Function inspecting empty interface interface{} / any using Type Switch
func describeValue(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type: int | Value x 2: %d\n", v*2)
	case string:
		fmt.Printf("Type: string | Length: %d\n", len(v))
	case Circle:
		fmt.Printf("Type: Circle | Radius: %.1f\n", v.Radius)
	default:
		fmt.Printf("Unknown Type: %T\n", v)
	}
}

func main() {
	fmt.Println("=== 7.1 Value vs Pointer Receivers ===")
	c := Circle{Radius: 5.0}
	fmt.Printf("Initial Circle Radius: %.1f | Area: %.2f\n", c.Radius, c.Area())

	c.Scale(2.0) // Modifies underlying struct
	fmt.Printf("Scaled Circle Radius: %.1f | New Area: %.2f\n", c.Radius, c.Area())

	fmt.Println("\n=== 7.2 Implicit Interface Implementation ===")
	rect := Rectangle{Width: 10.0, Height: 4.0}

	// Both Circle and Rectangle satisfy the Shape interface implicitly!
	printArea(c)
	printArea(rect)

	fmt.Println("\n=== 7.3 Type Assertions & Type Switches ===")
	// Type Assertion
	var val interface{} = "Hello Go"
	strVal, ok := val.(string)
	fmt.Printf("Type assertion successful: %t | Val: %s\n", ok, strVal)

	// Type Switch
	describeValue(100)
	describeValue("Golang Interfaces")
	describeValue(c)
	describeValue(3.14159)
}
