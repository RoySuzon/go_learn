# Complete Go (Golang) Programming Basic Course

Welcome to the **Complete Go (Golang) Programming Basic Course**. This guide covers foundational concepts, core language features, data structures, object-oriented concepts in Go, and concurrency with clear explanations and runnable code examples.

---

## Table of Contents
1. [Introduction & Program Structure](#1-introduction--program-structure)
   - 1.1 Hello World & Structure
   - 1.2 Packages and Imports
2. [Variables, Constants & Data Types](#2-variables-constants--data-types)
   - 2.1 Variable Declarations
   - 2.2 Basic Types & Zero Values
   - 2.3 Constants & `iota`
   - 2.4 Type Conversion
3. [Control Flow](#3-control-flow)
   - 3.1 Conditional Statements (`if`, `else if`, `else`)
   - 3.2 Switch Statements
   - 3.3 For Loops
4. [Functions & Error Handling](#4-functions--error-handling)
   - 4.1 Function Basics & Multiple Returns
   - 4.2 Named Returns & Variadic Functions
   - 4.3 Closures & Anonymous Functions
   - 4.4 Error Handling (`error`, `defer`, `panic`, `recover`)
5. [Data Structures](#5-data-structures)
   - 5.1 Arrays
   - 5.2 Slices
   - 5.3 Maps
   - 5.4 Structs & Composition
6. [Pointers](#6-pointers)
   - 6.1 Pointer Basics (`&` and `*`)
   - 6.2 Pass-by-Value vs Pass-by-Reference
7. [Methods & Interfaces](#7-methods--interfaces)
   - 7.1 Value Receivers vs Pointer Receivers
   - 7.2 Interfaces & Implicit Implementation
   - 7.3 Type Assertions & Type Switches
8. [Concurrency](#8-concurrency)
   - 8.1 Goroutines
   - 8.2 Channels (Unbuffered & Buffered)
   - 8.3 The `select` Statement
   - 8.4 `sync.WaitGroup` & `sync.Mutex`
9. [Modules & Standard Library](#9-modules--standard-library)
   - 9.1 Package Visibility Rules
   - 9.2 Essential Standard Packages (`fmt`, `strings`, `strconv`, `time`)

---

## 1. Introduction & Program Structure

### 1.1 Hello World & Structure
Every Go program starts with a `package` declaration. The `main` package with a `main()` function serves as the entry point of an executable application.

```go
// 01_hello.go
package main

import "fmt"

func main() {
    // Print string to standard output
    fmt.Println("Hello, World!")
}
```

### 1.2 Packages and Imports
Imports can be written as single lines or grouped in parentheses (factored import statement).

```go
// 02_imports.go
package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    fmt.Println("Current Time:", time.Now())
    fmt.Println("Square Root of 16:", math.Sqrt(16))
}
```

---

## 2. Variables, Constants & Data Types

### 2.1 Variable Declarations
Go offers several ways to declare variables: using `var` with explicit type, `var` with type inference, and short variable declaration (`:=`).

```go
// 03_variables.go
package main

import "fmt"

func main() {
    // Explicit type declaration
    var name string = "Alice"

    // Type inferred by compiler
    var age = 25

    // Short declaration syntax (only available inside functions)
    score := 95.5

    // Multiple variable declaration
    var x, y int = 10, 20
    city, country := "Dhaka", "Bangladesh"

    fmt.Printf("Name: %s, Age: %d, Score: %.1f\n", name, age, score)
    fmt.Printf("Location: %s, %s (x=%d, y=%d)\n", city, country, x, y)
}
```

### 2.2 Basic Types & Zero Values
Uninitialized variables in Go automatically receive their type's **zero value** (e.g., `0` for numbers, `""` for strings, `false` for booleans, `nil` for pointers/slices/maps).

```go
// 04_zero_values.go
package main

import "fmt"

func main() {
    var i int       // Zero value: 0
    var f float64   // Zero value: 0.0
    var b bool      // Zero value: false
    var s string    // Zero value: ""

    fmt.Printf("int zero value: %v\n", i)
    fmt.Printf("float zero value: %v\n", f)
    fmt.Printf("bool zero value: %v\n", b)
    fmt.Printf("string zero value: %q\n", s)
}
```

### 2.3 Constants & `iota`
Constants are declared using `const`. `iota` is a special constant generator used to create successive untyped integer constants.

```go
// 05_constants.go
package main

import "fmt"

const ApplicationName = "GoCourse"

// Days of the week using iota
const (
    Sunday = iota // 0
    Monday        // 1
    Tuesday       // 2
    Wednesday     // 3
    Thursday      // 4
    Friday        // 5
    Saturday      // 6
)

func main() {
    fmt.Println("AppName:", ApplicationName)
    fmt.Println("Sunday index:", Sunday)
    fmt.Println("Wednesday index:", Wednesday)
    fmt.Println("Friday index:", Friday)
}
```

### 2.4 Type Conversion
Go requires explicit type conversion between different types.

```go
// 06_type_conversion.go
package main

import "fmt"

func main() {
    var a int = 42
    var b float64 = float64(a)
    var c uint = uint(b)

    fmt.Printf("a: %T = %v\n", a, a)
    fmt.Printf("b: %T = %v\n", b, b)
    fmt.Printf("c: %T = %v\n", c, c)
}
```

---

## 3. Control Flow

### 3.1 Conditional Statements (`if`, `else if`, `else`)
Go supports optional short initialization statements before the boolean condition in `if`.

```go
// 07_if_else.go
package main

import "fmt"

func main() {
    num := 15

    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // If with short statement (limit is scoped only inside if-else block)
    if limit := 20; num < limit {
        fmt.Printf("%d is within limit %d\n", num, limit)
    }
}
```

### 3.2 Switch Statements
In Go, `break` is automatic at the end of each case. You can switch on values, expressions, or types.

```go
// 08_switch.go
package main

import (
    "fmt"
    "time"
)

func main() {
    day := "Monday"

    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("Weekday")
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Invalid day")
    }

    // Tagless switch (behaves like if-else chain)
    t := time.Now().Hour()
    switch {
    case t < 12:
        fmt.Println("Good Morning!")
    case t < 17:
        fmt.Println("Good Afternoon!")
    default:
        fmt.Println("Good Evening!")
    }
}
```

### 3.3 For Loops
`for` is Go's only looping construct. It can act as a standard loop, a while loop, or an infinite loop.

```go
// 09_loops.go
package main

import "fmt"

func main() {
    // 1. Standard 3-component loop
    fmt.Print("Standard loop: ")
    for i := 1; i <= 3; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 2. "While"-style loop
    fmt.Print("While-style loop: ")
    n := 1
    for n <= 3 {
        fmt.Printf("%d ", n)
        n++
    }
    fmt.Println()

    // 3. Infinite loop with break
    fmt.Print("Infinite loop with break: ")
    count := 0
    for {
        count++
        if count > 3 {
            break
        }
        fmt.Printf("%d ", count)
    }
    fmt.Println()
}
```

---

## 4. Functions & Error Handling

### 4.1 Function Basics & Multiple Returns
Go functions can return multiple values, typically used to return `(result, error)`.

```go
// 10_functions.go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

// Function returning multiple values
func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "division by zero error"
    }
    return a / b, ""
}

func main() {
    sum := add(10, 20)
    fmt.Println("Sum:", sum)

    res, err := divide(10, 2)
    if err == "" {
        fmt.Println("Result:", res)
    }

    _, errZero := divide(10, 0)
    fmt.Println("Error:", errZero)
}
```

### 4.2 Named Returns & Variadic Functions
Named return variables are initialized to their zero values. Variadic functions accept zero or more trailing arguments.

```go
// 11_variadic_named.go
package main

import "fmt"

// Named return values
func calculate(length, width int) (area int, perimeter int) {
    area = length * width
    perimeter = 2 * (length + width)
    return // Naked return
}

// Variadic function
func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    a, p := calculate(5, 3)
    fmt.Printf("Area: %d, Perimeter: %d\n", a, p)

    fmt.Println("Sum of numbers:", sumAll(1, 2, 3, 4, 5))
}
```

### 4.3 Closures & Anonymous Functions
Go supports anonymous functions and closures (functions that reference variables outside their body).

```go
// 12_closures.go
package main

import "fmt"

func sequenceGenerator() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    next := sequenceGenerator()

    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
    fmt.Println(next()) // 3

    // New sequence instance
    next2 := sequenceGenerator()
    fmt.Println(next2()) // 1
}
```

### 4.4 Error Handling (`error`, `defer`, `panic`, `recover`)
Go handles errors explicitly using the `error` type. `defer` schedules a function call to run when the surrounding function returns.

```go
// 13_errors_defer.go
package main

import (
    "errors"
    "fmt"
)

func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func demonstratePanicRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Triggering panic...")
    panic("critical system failure!")
}

func main() {
    // Defer executes in Last-In-First-Out (LIFO) order upon main return
    defer fmt.Println("Deferred message 1 (executed at end)")
    defer fmt.Println("Deferred message 2 (executed before msg 1)")

    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("Handled Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    demonstratePanicRecover()
}
```

---

## 5. Data Structures

### 5.1 Arrays
Arrays in Go have a fixed size that is part of their type.

```go
// 14_arrays.go
package main

import "fmt"

func main() {
    // Array of 5 integers
    var numbers [5]int = [5]int{10, 20, 30, 40, 50}
    
    // Auto-length detection
    names := [...]string{"Alice", "Bob", "Charlie"}

    fmt.Println("Numbers array:", numbers)
    fmt.Println("Array Length:", len(numbers))
    fmt.Println("Names array:", names)
}
```

### 5.2 Slices
Slices are dynamically sized, flexible views into arrays. Created using `make()` or literal syntax.

```go
// 15_slices.go
package main

import "fmt"

func main() {
    // Slice declaration
    fruits := []string{"Apple", "Banana", "Cherry"}

    // Append items
    fruits = append(fruits, "Date", "Elderberry")
    fmt.Println("Fruits slice:", fruits)

    // Slice expression: fruits[low:high]
    subSlice := fruits[1:4] // Elements from index 1 up to 3
    fmt.Println("Sub-slice [1:4]:", subSlice)

    // Creating slice with make(type, len, cap)
    numbers := make([]int, 3, 5) // length 3, capacity 5
    numbers[0] = 100
    numbers[1] = 200
    numbers[2] = 300
    fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))
}
```

### 5.3 Maps
Maps are hash tables storing key-value pairs.

```go
// 16_maps.go
package main

import "fmt"

func main() {
    // Initialize map
    userAges := map[string]int{
        "Alice": 28,
        "Bob":   34,
    }

    // Adding/updating keys
    userAges["Charlie"] = 22

    // Retrieve value and check existence (comma-ok idiom)
    age, exists := userAges["Bob"]
    fmt.Printf("Bob's age: %d (Exists: %t)\n", age, exists)

    val, ok := userAges["David"]
    fmt.Printf("David's age: %d (Exists: %t)\n", val, ok)

    // Delete key
    delete(userAges, "Alice")

    // Iterate over map
    fmt.Println("Remaining users:")
    for key, val := range userAges {
        fmt.Printf("- %s: %d years old\n", key, val)
    }
}
```

### 5.4 Structs & Composition
Structs are typed collections of fields. Go uses composition (struct embedding) instead of traditional OOP inheritance.

```go
// 17_structs.go
package main

import "fmt"

type Address struct {
    City    string
    Country string
}

type Employee struct {
    ID      int
    Name    string
    Salary  float64
    Address // Embedded struct (Composition)
}

func main() {
    emp := Employee{
        ID:     101,
        Name:   "Sarah",
        Salary: 75000.50,
        Address: Address{
            City:    "New York",
            Country: "USA",
        },
    }

    fmt.Printf("Employee: %+v\n", emp)
    // Promoted fields access
    fmt.Printf("Name: %s, City: %s\n", emp.Name, emp.City)
}
```

---

## 6. Pointers

### 6.1 Pointer Basics (`&` and `*`)
Pointers store the memory address of a value.

```go
// 18_pointers.go
package main

import "fmt"

func main() {
    val := 42
    var ptr *int = &val // & gets address of val

    fmt.Println("Value of val:", val)
    fmt.Println("Address of val (&val):", ptr)
    fmt.Println("Dereferenced ptr (*ptr):", *ptr) // * gets value at address

    // Modify original value via pointer
    *ptr = 100
    fmt.Println("New value of val:", val)
}
```

### 6.2 Pass-by-Value vs Pass-by-Reference
Go is strictly **pass-by-value**. To modify caller variables, pass a pointer address.

```go
// 19_pass_by_pointer.go
package main

import "fmt"

func updateValueByValue(x int) {
    x = 500 // Modifies local copy only
}

func updateValueByPointer(x *int) {
    *x = 500 // Modifies underlying variable
}

func main() {
    number := 50

    updateValueByValue(number)
    fmt.Println("After updateValueByValue:", number) // 50

    updateValueByPointer(&number)
    fmt.Println("After updateValueByPointer:", number) // 500
}
```

---

## 7. Methods & Interfaces

### 7.1 Value Receivers vs Pointer Receivers
Methods attach functions to types. Pointer receivers allow modifying struct fields.

```go
// 20_methods.go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

// Value receiver (does not modify original struct)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver (can modify struct contents)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Initial Area:", rect.Area())

    rect.Scale(2.0)
    fmt.Printf("Scaled Dimensions: %.1f x %.1f\n", rect.Width, rect.Height)
    fmt.Println("New Area:", rect.Area())
}
```

### 7.2 Interfaces & Implicit Implementation
Interfaces specify behavior. Any type implementing required methods implicitly satisfies the interface.

```go
// 21_interfaces.go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f | Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    c := Circle{Radius: 5}
    fmt.Print("Circle Info -> ")
    printShapeInfo(c)
}
```

### 7.3 Type Assertions & Type Switches
Inspect the concrete dynamic type of an interface variable.

```go
// 22_type_switch.go
package main

import "fmt"

func describeType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    case bool:
        fmt.Println("Boolean:", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func main() {
    describeType(42)
    describeType("Golang")
    describeType(true)
    describeType(3.14)
}
```

---

## 8. Concurrency

### 8.1 Goroutines
Goroutines are lightweight threads managed by the Go runtime. Started with the `go` keyword.

```go
// 23_goroutines.go
package main

import (
    "fmt"
    "time"
)

func printMessage(text string) {
    for i := 1; i <= 3; i++ {
        fmt.Printf("%s: %d\n", text, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Run in background goroutine
    go printMessage("Async Task")

    // Run in main goroutine
    printMessage("Main Task")
}
```

### 8.2 Channels (Unbuffered & Buffered)
Channels allow goroutines to communicate and synchronize execution.

```go
// 24_channels.go
package main

import "fmt"

func worker(ch chan string) {
    ch <- "Task completed!" // Send data to channel
}

func main() {
    // 1. Unbuffered Channel
    ch := make(chan string)
    go worker(ch)

    msg := <-ch // Receive data from channel
    fmt.Println("Received from unbuffered channel:", msg)

    // 2. Buffered Channel (capacity 2)
    bufCh := make(chan int, 2)
    bufCh <- 100
    bufCh <- 200

    fmt.Println("Buffered val 1:", <-bufCh)
    fmt.Println("Buffered val 2:", <-bufCh)
}
```

### 8.3 The `select` Statement
`select` lets a goroutine wait on multiple channel communication operations.

```go
// 25_select.go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "Fast result"
    }()

    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "Slow result"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received:", msg2)
        }
    }
}
```

### 8.4 `sync.WaitGroup` & `sync.Mutex`
`sync.WaitGroup` waits for a collection of goroutines to finish. `sync.Mutex` ensures mutual exclusion to prevent data races.

```go
// 26_sync.go
package main

import (
    "fmt"
    "sync"
)

type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
    defer wg.Done()
    
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}

func main() {
    var wg sync.WaitGroup
    counter := SafeCounter{}

    // Launch 10 goroutines concurrently
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go counter.Increment(&wg)
    }

    wg.Wait() // Block until all 10 goroutines complete
    fmt.Println("Final Counter Value:", counter.value)
}
```

---

## 9. Modules & Standard Library

### 9.1 Package Visibility Rules
- **Exported Identifiers**: Capitalized first letter (e.g., `fmt.Println`, `MyStruct`, `ExportedFunction`). Visible outside the package.
- **Unexported Identifiers**: Lowercase first letter (e.g., `privateVar`, `internalHelper`). Accessible only within the declaring package.

Initialize a module in your terminal using:
```bash
go mod init myproject
```

### 9.2 Essential Standard Packages
Below are examples of working with common Go standard library packages.

```go
// 27_stdlib_helpers.go
package main

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

func main() {
    // 1. Strings Package
    text := "  Hello Go World!  "
    trimmed := strings.TrimSpace(text)
    fmt.Println("Upper:", strings.ToUpper(trimmed))
    fmt.Println("Contains 'Go':", strings.Contains(trimmed, "Go"))

    // 2. Strconv Package
    strNum := "123"
    num, _ := strconv.Atoi(strNum)
    fmt.Println("Parsed Number + 10:", num+10)

    // 3. Time Package
    now := time.Now()
    fmt.Println("Formatted Date:", now.Format("2006-01-02 15:04:05"))
}
```

---

## Conclusion & Next Steps

Congratulations! You have completed the basic Go programming course outline. 

### Practice Projects to Build Next:
1. **CLI Todo App**: Create a command-line application using `os.Args`, `bufio.Scanner`, and JSON file storage.
2. **REST API**: Build a simple JSON API server using `net/http`.
3. **Concurrent Web Crawler**: Fetch multiple URLs concurrently using Goroutines, Channels, and `sync.WaitGroup`.
