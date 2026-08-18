package main

import "fmt"

// Embedded struct for composition
type Address struct {
	City    string
	Country string
}

// Parent struct embedding Address
type User struct {
	ID      int
	Name    string
	Email   string
	Address // Struct embedding (Composition)
}

func main() {
	fmt.Println("=== 5.1 Arrays ===")
	// Fixed size arrays (size is part of the array's type)
	var scores [3]int = [3]int{90, 85, 95}
	colors := [...]string{"Red", "Green", "Blue"} // Auto-calculated length

	fmt.Println("Scores array:", scores)
	fmt.Println("Colors array:", colors, "| Length:", len(colors))

	fmt.Println("\n=== 5.2 Slices ===")
	// Slices are dynamically sized views over underlying arrays
	languages := []string{"Go", "Python", "JavaScript"}

	// Appending items dynamically
	languages = append(languages, "Rust", "C++")
	fmt.Println("Languages Slice:", languages)

	// Slice expression [low:high]
	topTwo := languages[:2] // Indices 0 and 1
	fmt.Println("Top 2 Languages:", topTwo)

	// Creating slices with make(type, len, cap)
	buf := make([]int, 3, 5) // length=3, capacity=5
	buf[0], buf[1], buf[2] = 10, 20, 30
	fmt.Printf("Slice: %v | Len: %d | Cap: %d\n", buf, len(buf), cap(buf))

	fmt.Println("\n=== 5.3 Maps ===")
	// Key-value pairs
	userRoles := map[string]string{
		"goutom": "Admin",
		"alice":  "Developer",
	}

	// Insert/Update
	userRoles["bob"] = "Tester"

	// Key existence check (comma-ok idiom)
	role, exists := userRoles["goutom"]
	fmt.Printf("User 'goutom' -> Role: %s (Exists: %t)\n", role, exists)

	missingRole, found := userRoles["charlie"]
	fmt.Printf("User 'charlie' -> Role: %s (Exists: %t)\n", missingRole, found)

	// Delete key
	delete(userRoles, "alice")

	// Iterate map
	fmt.Println("Current User Roles:")
	for username, r := range userRoles {
		fmt.Printf("  - %s: %s\n", username, r)
	}

	fmt.Println("\n=== 5.4 Structs & Composition ===")
	u := User{
		ID:    101,
		Name:  "Goutom Roy",
		Email: "goutom@example.com",
		Address: Address{
			City:    "Dhaka",
			Country: "Bangladesh",
		},
	}

	fmt.Printf("User Struct: %+v\n", u)
	// Promoted fields from embedded struct Address
	fmt.Printf("User Name: %s | City: %s | Country: %s\n", u.Name, u.City, u.Country)
}
