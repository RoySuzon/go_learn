package main

import "fmt"

type Account struct {
	Balance float64
}

// Function taking argument by VALUE (makes a local copy)
func updateByValue(val int) {
	val = 999 // Only changes local copy
}

// Function taking argument by POINTER (modifies original variable)
func updateByPointer(ptr *int) {
	*ptr = 999 // Dereferences pointer and updates memory location
}

// Struct modification using Pointer
func deposit(acc *Account, amount float64) {
	acc.Balance += amount // Go automatically dereferences pointers to structs!
}

func main() {
	fmt.Println("=== 6.1 Pointer Operators (& and *) ===")
	num := 42

	// & gets the memory address of num
	var ptr *int = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Memory address of num (&num):", ptr)
	fmt.Println("Dereferenced value (*ptr):", *ptr)

	// Modifying original value using pointer dereferencing
	*ptr = 100
	fmt.Println("New value of num after *ptr = 100:", num)

	fmt.Println("\n=== 6.2 Pass-by-Value vs Pass-by-Pointer ===")
	val := 50

	updateByValue(val)
	fmt.Println("After updateByValue:", val) // Remains 50

	updateByPointer(&val)
	fmt.Println("After updateByPointer:", val) // Becomes 999

	fmt.Println("\n=== 6.3 Modifying Structs via Pointers ===")
	myAcc := Account{Balance: 250.00}
	fmt.Printf("Initial Balance: $%.2f\n", myAcc.Balance)

	deposit(&myAcc, 150.50)
	fmt.Printf("Updated Balance after Deposit: $%.2f\n", myAcc.Balance)
}
