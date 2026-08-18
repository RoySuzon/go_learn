package main

import "fmt"

// ---------------------------------------------------------
// 1. GENERIC FUNCTION (Any type T)
// ---------------------------------------------------------

func PrintSlice[T any](items []T) {
	fmt.Print("[ ")
	for _, item := range items {
		fmt.Printf("%v ", item)
	}
	fmt.Println("]")
}

// ---------------------------------------------------------
// 2. GENERIC CONSTRAINTS (Ordered types for comparison)
// ---------------------------------------------------------

type Number interface {
	~int | ~int64 | ~float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ---------------------------------------------------------
// 3. GENERIC DATA STRUCTURE (Stack[T])
// ---------------------------------------------------------

type GenericStack[T any] struct {
	items []T
}

func (s *GenericStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *GenericStack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	topIndex := len(s.items) - 1
	item := s.items[topIndex]
	s.items = s.items[:topIndex]
	return item, true
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 Lesson 11: Generics in Go (Go 1.18+ টাইপ প্যারামিটার)")
	fmt.Println("==================================================")

	// 1. Generic Print
	intSlice := []int{10, 20, 30}
	stringSlice := []string{"Go", "Generics", "Powerful"}

	fmt.Print("১. ইন্টিজার স্লাইস: ")
	PrintSlice(intSlice)

	fmt.Print("২. স্ট্রিং স্লাইস: ")
	PrintSlice(stringSlice)

	// 2. Generic Max
	fmt.Printf("\n৩. Max(15, 42) = %d\n", Max(15, 42))
	fmt.Printf("৪. Max(99.5, 23.4) = %.2f\n", Max(99.5, 23.4))

	// 3. Generic Stack
	fmt.Println("\n৫. Generic Stack[T] পরীক্ষা:")
	intStack := &GenericStack[int]{}
	intStack.Push(100)
	intStack.Push(200)
	val, _ := intStack.Pop()
	fmt.Println("   - intStack Popped:", val)

	strStack := &GenericStack[string]{}
	strStack.Push("Hello")
	strStack.Push("Generics")
	strVal, _ := strStack.Pop()
	fmt.Println("   - strStack Popped:", strVal)
}
