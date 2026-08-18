package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 9.1 Package Export Rules:
// Capitalized identifiers (ExportedStruct, ExportedFunc) are PUBLIC and accessible across packages.
// Lowercase identifiers (internalStruct, internalFunc) are PRIVATE to the declaring package.

type ExportedConfig struct {
	AppName string // Public field
	port    int    // Private field
}

func main() {
	fmt.Println("=== 9.1 Package Visibility Rules ===")
	cfg := ExportedConfig{
		AppName: "Go Master App",
		port:    8080, // Accessible inside main package
	}
	fmt.Printf("Config -> AppName: %s | Internal Port: %d\n", cfg.AppName, cfg.port)

	fmt.Println("\n=== 9.2 Standard Package: strings ===")
	str := "  Welcome to Go Programming!  "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed String:", trimmed)
	fmt.Println("Uppercase:", strings.ToUpper(trimmed))
	fmt.Println("Contains 'Go':", strings.Contains(trimmed, "Go"))
	fmt.Println("Replace 'Go' with 'Golang':", strings.ReplaceAll(trimmed, "Go", "Golang"))

	fmt.Println("\n=== 9.3 Standard Package: strconv ===")
	// Convert String to Int
	numStr := "1250"
	num, err := strconv.Atoi(numStr)
	if err == nil {
		fmt.Printf("Parsed Number + 250: %d\n", num+250)
	}

	// Convert Int to String
	strConverted := strconv.Itoa(500)
	fmt.Printf("Converted String: %s (Type: %T)\n", strConverted, strConverted)

	fmt.Println("\n=== 9.4 Standard Package: time ===")
	now := time.Now()
	// Go's unique time formatting reference date: Mon Jan 2 15:04:05 MST 2006 (1 2 3 4 5 6)
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Date/Time:", formattedTime)

	fmt.Println("\n=== 9.5 Standard Package: os ===")
	hostname, _ := os.Hostname()
	fmt.Println("System Hostname:", hostname)
	fmt.Println("Current Process ID (PID):", os.Getpid())
}
