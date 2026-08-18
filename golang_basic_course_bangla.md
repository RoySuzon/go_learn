# সম্পূর্ণ গো (Golang) প্রোগ্রামিং বেসিক কোর্স (Bangla Guide)

স্বাগতম **সম্পূর্ণ গো (Golang) প্রোগ্রামিং বেসিক কোর্স**-এ। এই নির্দেশিকায় Go প্রোগ্রামিং ভাষার মূল ধারণাসমূহ, সিনট্যাক্স, ডেটা স্ট্রাকচার, অবজেক্ট-ওরিয়েন্টেড ধাঁচ, এরর হ্যান্ডলিং এবং কনকারেন্সি সহজ বাংলায় বিস্তারিত ব্যাখ্যা এবং নির্ভুল কোড উদাহরণসহ তুলে ধরা হয়েছে।

---

## সূচিপত্র (Table of Contents)
1. [সূচনা ও প্রোগ্রাম স্ট্রাকচার (Introduction & Program Structure)](#1-সূচনা-ও-প্রোগ্রাম-স্ট্রাকচার)
   - 1.1 হ্যালো ওয়ার্ল্ড ও প্রোগ্রামের গঠন
   - 1.2 প্যাকেজ এবং ইমপোর্ট (Packages & Imports)
2. [ভ্যারিয়েবল, কনস্ট্যান্ট এবং ডেটা টাইপ (Variables, Constants & Data Types)](#2-ভ্যারিয়েবল-কনস্ট্যান্ট-এবং-ডেটা-টাইপ)
   - 2.1 ভ্যারিয়েবল ডিক্লেয়ারেশন (Variable Declarations)
   - 2.2 ডিফল্ট মান বা জিরো ভ্যালু (Zero Values)
   - 2.3 কনস্ট্যান্ট এবং `iota` (Constants & `iota`)
   - 2.4 টাইপ কনভার্সন (Type Conversion)
3. [কন্ট্রোল ফ্লো (Control Flow)](#3-কন্ট্রোল-ফ্লো)
   - 3.1 কন্ডিশনাল স্টেটমেন্ট (`if`, `else if`, `else`)
   - 3.2 সুইচ স্টেটমেন্ট (`switch`)
   - 3.3 ফর লুপ (`for` loops)
4. [ফাংশন এবং এরর হ্যান্ডলিং (Functions & Error Handling)](#4-ফাংশন-এবং-এরর-হ্যান্ডলিং)
   - 4.1 ফাংশন ও মাল্টিপল রিটার্ন ভ্যালু
   - 4.2 নেমড রিটার্ন এবং ভ্যারিয়াডিক ফাংশন
   - 4.3 ক্লোজার এবং অ্যানোনিমাস ফাংশন (Closures)
   - 4.4 এরর হ্যান্ডলিং (`error`, `defer`, `panic`, `recover`)
5. [ডেটা স্ট্রাকচার (Data Structures)](#5-ডেটা-স্ট্রাকচার)
   - 5.1 অ্যারে (Arrays)
   - 5.2 স্লাইস (Slices)
   - 5.3 ম্যাপ (Maps)
   - 5.4 স্ট্রাক্ট এবং কম্পোজিশন (Structs & Composition)
6. [পয়েন্টার (Pointers)](#6-পয়েন্টার)
   - 6.1 পয়েন্টার বেসিক্স (`&` এবং `*`)
   - 6.2 পাশ-বাই-ভ্যালু বনাম পাশ-বাই-পয়েন্টার
7. [মেথড এবং ইন্টারফেস (Methods & Interfaces)](#7-মেথড-এবং-ইন্টারফেস)
   - 7.1 ভ্যালু রিসিভার বনাম পয়েন্টার রিসিভার
   - 7.2 ইন্টারফেস এবং ইমপ্লিসিট ইমপ্লিমেন্টেশন
   - 7.3 টাইপ অ্যাসারশন এবং টাইপ সুইচ
8. [কনকারেন্সি (Concurrency)](#8-কনকারেন্সি)
   - 8.1 গোরুটিন (Goroutines)
   - 8.2 চ্যানেল (Channels)
   - 8.3 সিলেক্ট স্টেটমেন্ট (`select`)
   - 8.4 `sync.WaitGroup` এবং `sync.Mutex`
9. [মডিউল এবং স্ট্যান্ডার্ড লাইব্রেরি (Modules & Standard Library)](#9-মডিউল-এবং-স্ট্যান্ডার্ড-লাইব্রেরি)
   - 9.1 প্যাকেজ ভিজিবিলিটি রুলস (Public/Private Rules)
   - 9.2 প্রয়োজনীয় স্ট্যান্ডার্ড প্যাকেজসমূহ (`fmt`, `strings`, `strconv`, `time`, `os`)

---

## 1. সূচনা ও প্রোগ্রাম স্ট্রাকচার

### 1.1 হ্যালো ওয়ার্ল্ড ও প্রোগ্রামের গঠন
প্রতিটি রানঅ্যাবল (Executable) গো প্রোগ্রাম অবশ্যই `package main` দিয়ে শুরু হতে হয় এবং এতে একটি `main()` ফাংশন থাকতে হয়।

```go
// 01_hello.go
package main

import "fmt" // স্ট্যান্ডার্ড ইনপুট-আউটপুট প্যাকেজ

func main() {
    // কনসোলে টেক্সট প্রিন্ট করার জন্য fmt.Println ব্যবহৃত হয়
    fmt.Println("হ্যালো ওয়ার্ল্ড! গো প্রোগ্রামিং-এ স্বাগতম।")
}
```
**ব্যাখ্যা:**
- `package main`: এই ফাইলটি একটি এক্সিকিউটেবল প্রোগ্রাম হিসেবে রান হবে।
- `import "fmt"`: ফরম্যাটেড ইনপুট/আউটপুটের জন্য স্ট্যান্ডার্ড প্যাকেজটি যুক্ত করা হয়েছে।
- `func main()`: প্রোগ্রামের মূল এন্ট্রি পয়েন্ট।

---

### 1.2 প্যাকেজ এবং ইমপোর্ট
একাধিক প্যাকেজ ইমপোর্ট করার জন্য ব্র্যাকেট `import (...)` ব্যবহার করা হয়।

```go
// 02_imports.go
package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    fmt.Println("বর্তমান সময়:", time.Now().Format("2006-01-02 15:04:05"))
    fmt.Println("১৬-এর বর্গমূল:", math.Sqrt(16))
}
```

---

## 2. ভ্যারিয়েবল, কনস্ট্যান্ট এবং ডেটা টাইপ

### 2.1 ভ্যারিয়েবল ডিক্লেয়ারেশন
গো-তে ভ্যারিয়েবল ঘোষণার ৩টি প্রধান উপায় রয়েছে:

```go
// 03_variables.go
package main

import "fmt"

func main() {
    // ১. স্পষ্ট টাইপসহ ঘোষণা (Explicit type declaration)
    var name string = "গৌতম"

    // ২. টাইপ অনুমান (Type inference)
    var age = 26

    // ৩. শর্ট ভ্যারিয়েবল ডিক্লেয়ারেশন syntax := (শুধুমাত্র ফাংশনের ভেতরে ব্যবহারযোগ্য)
    score := 95.5

    // একাধিক ভ্যারিয়েবল একসাথে ঘোষণা
    x, y := 10, 20

    fmt.Printf("নাম: %s, বয়স: %d, স্কোর: %.1f\n", name, age, score)
    fmt.Printf("x = %d, y = %d\n", x, y)
}
```

---

### 2.2 ডিফল্ট মান বা জিরো ভ্যালু (Zero Values)
কোনো ভ্যারিয়েবলে মান না দিলে Go স্বয়ংক্রিয়ভাবে তার টাইপ অনুযায়ী ডিফল্ট জিরো ভ্যালু সেট করে।

```go
// 04_zero_values.go
package main

import "fmt"

func main() {
    var i int       // ডিফল্ট মান: 0
    var f float64   // ডিফল্ট মান: 0.0
    var b bool      // ডিফল্ট মান: false
    var s string    // ডিফল্ট মান: "" (খালি স্ট্রিং)

    fmt.Printf("int: %v | float: %v | bool: %v | string: %q\n", i, f, b, s)
}
```

---

### 2.3 কনস্ট্যান্ট এবং `iota`
`const` দিয়ে এমন ভ্যালু ঘোষণা করা হয় যা প্রোগ্রামে কখনো পরিবর্তন হবে না। `iota` পরপর ক্রমিক সংখ্যা তৈরি করতে ব্যবহৃত হয়।

```go
// 05_constants.go
package main

import "fmt"

const AppName = "GoCourseBangla"

// iota ব্যবহার করে অটো-ইনক্রিমেন্ট কনস্ট্যান্ট তৈরি
const (
    Pending  = iota // 0
    Approved        // 1
    Rejected        // 2
)

func main() {
    fmt.Println("অ্যাপ নাম:", AppName)
    fmt.Printf("স্ট্যাটাস কোড -> পেন্ডিং: %d, অ্যাপ্রুভড: %d, রিজেক্টেড: %d\n", Pending, Approved, Rejected)
}
```

---

### 2.4 টাইপ কনভার্সন
Go-তে কোনো অটোমেটিক (Implicit) টাইপ পরিবর্তন হয় না। এক টাইপ থেকে অন্য টাইপে রূপান্তর করতে টাইপ কাস্টিং করতে হয়।

```go
// 06_type_conversion.go
package main

import "fmt"

func main() {
    var a int = 42
    var b float64 = float64(a) // int থেকে float64 এ স্পষ্ট পরিবর্তন
    var c uint = uint(b)

    fmt.Printf("a: %T = %v\n", a, a)
    fmt.Printf("b: %T = %v\n", b, b)
    fmt.Printf("c: %T = %v\n", c, c)
}
```

---

## 3. কন্ট্রোল ফ্লো

### 3.1 কন্ডিশনাল স্টেটমেন্ট (`if`, `else if`, `else`)
Go-তে `if`-এর কন্ডিশনের আগে একটি শর্ট স্টেটমেন্ট লেখা যায়।

```go
// 07_if_else.go
package main

import "fmt"

func main() {
    number := 15

    if number%2 == 0 {
        fmt.Println(number, "হলো জোড় সংখ্যা")
    } else {
        fmt.Println(number, "হলো বিজোড় সংখ্যা")
    }

    // শর্ট স্টেটমেন্টসহ if (limit ভ্যারিয়েবলটি কেবল এই if-else ব্লকের ভেতরেই কার্যকর)
    if limit := 20; number < limit {
        fmt.Printf("%d সংখ্যাটি সীমা %d-এর নিচে\n", number, limit)
    }
}
```

---

### 3.2 সুইচ স্টেটমেন্ট
Go-এর `switch`-এ প্রতিটি কেসের শেষে স্বয়ংক্রিয়ভাবে `break` হয়ে যায় (আলাদা করে break লিখতে হয় না)।

```go
// 08_switch.go
package main

import "fmt"

func main() {
    day := "Wednesday"

    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println(day, "হলো কর্মদিবস (Weekday)")
    case "Saturday", "Sunday":
        fmt.Println(day, "হলো ছুটির দিন (Weekend)")
    default:
        fmt.Println("অবৈধ দিন")
    }
}
```

---

### 3.3 ফর লুপ (`for` loops)
Go-তে একমাত্র লুপ কিওয়ার্ড হলো `for` (কোনো while বা do-while নেই)।

```go
// 09_loops.go
package main

import "fmt"

func main() {
    // ১. সাধারণ ৩-ধাপের ফর লুপ
    fmt.Print("ফর লুপ: ")
    for i := 1; i <= 3; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // ২. While-ধাঁচের লুপ
    count := 3
    fmt.Print("While-ধাঁচের লুপ: ")
    for count > 0 {
        fmt.Printf("%d ", count)
        count--
    }
    fmt.Println()

    // ৩. ইনফিনিট লুপ এবং break
    n := 0
    fmt.Print("ইনফিনিট লুপ: ")
    for {
        n++
        if n > 3 {
            break // লুপ থেকে বের হয়ে যাবে
        }
        fmt.Printf("%d ", n)
    }
    fmt.Println()
}
```

---

## 4. ফাংশন এবং এরর হ্যান্ডলিং

### 4.1 ফাংশন ও মাল্টিপল রিটার্ন ভ্যালু
Go-এর ফাংশন থেকে একাধিক মান (সাধারণত রেজাল্ট এবং এরর) রিটার্ন করা যায়।

```go
// 10_functions.go
package main

import (
    "errors"
    "fmt"
)

// ভাগ করার ফাংশন যা ফলাফল এবং এরর দুটোই রিটার্ন করে
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("শূন্য (0) দিয়ে ভাগ করা সম্ভব নয়")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("এরর:", err)
    } else {
        fmt.Println("ভাগফল:", result)
    }

    _, errZero := divide(10, 0)
    if errZero != nil {
        fmt.Println("হ্যান্ডেলকৃত এরর:", errZero)
    }
}
```

---

### 4.2 নেমড রিটার্ন এবং ভ্যারিয়াডিক ফাংশন

```go
// 11_variadic_named.go
package main

import "fmt"

// নেমড রিটার্ন ভ্যালু (Named return values)
func getRectangleStats(l, w float64) (area float64, perimeter float64) {
    area = l * w
    perimeter = 2 * (l + w)
    return // Naked return
}

// ভ্যারিয়াডিক ফাংশন (অননির্দিষ্ট সংখ্যক আর্গুমেন্ট গ্রহণ করে)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    a, p := getRectangleStats(5, 3)
    fmt.Printf("ক্ষেত্রফল: %.1f, পরিসীমা: %.1f\n", a, p)

    fmt.Println("মোট যোগফল:", sum(10, 20, 30, 40))
}
```

---

### 4.3 ক্লোজার (Closures)
একটি ক্লোজার হলো এমন একটি অ্যানোনিমাস ফাংশন যা তার বাইরের ভ্যারিয়েবলগুলোর রেফারেন্স ধরে রাখতে পারে।

```go
// 12_closures.go
package main

import "fmt"

func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := makeCounter()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
```

---

### 4.4 এরর হ্যান্ডলিং (`defer`, `panic`, `recover`)
- `defer`: ফাংশনের কাজ শেষ হওয়া পর্যন্ত কোনো স্টেটমেন্টকে স্থগিত রাখে (LIFO নিয়মে চলে)।
- `panic`: প্রোগ্রাম থামিয়ে দেয়।
- `recover`: প্যানিক হওয়া প্রোগ্রামকে ক্র্যাশ হওয়া থেকে রক্ষা করে।

```go
// 13_defer_panic_recover.go
package main

import "fmt"

func safeExecution() {
    // recover সর্বদা defer ব্লকের ভেতরেই লিখতে হয়
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("--> প্যানিক থেকে রিকভার করা হয়েছে:", r)
        }
    }()

    fmt.Println("প্রোগ্রাম চলছে...")
    panic("গুরুতর সমস্যা দেখা দিয়েছে!") // প্যানিক ট্রিগার
}

func main() {
    defer fmt.Println("--> ডিফার্ড ক্লিনআপ বার্তা (সবার শেষে রান হবে)")

    safeExecution()
}
```

---

## 5. ডেটা স্ট্রাকচার

### 5.1 অ্যারে (Arrays)
অ্যারের আকার (Size) নির্দিষ্ট থাকে এবং তা এর টাইপের অংশ।

```go
// 14_arrays.go
package main

import "fmt"

func main() {
    var numbers [3]int = [3]int{10, 20, 30}
    fmt.Println("অ্যাের উপাদান:", numbers)
    fmt.Println("অ্যারের দৈর্ঘ্য:", len(numbers))
}
```

---

### 5.2 স্লাইস (Slices)
স্লাইস হলো পরিবর্তনশীল এবং গতিশীল (Dynamic) ডেটা লিস্ট।

```go
// 15_slices.go
package main

import "fmt"

func main() {
    // স্লাইস তৈরি
    fruits := []string{"আম", "কলা"}

    // নতুন উপাদান যুক্তকরণ (append)
    fruits = append(fruits, "লিচু", "কাঁঠাল")
    fmt.Println("ফলসমূহ:", fruits)

    // make(type, len, cap) দিয়ে স্লাইস তৈরি
    buf := make([]int, 2, 5) // দৈর্ঘ্য ২, ক্যাপাসিটি ৫
    buf[0] = 100
    buf[1] = 200
    fmt.Printf("স্লাইস: %v | দৈর্ঘ্য: %d | ক্যাপাসিটি: %d\n", buf, len(buf), cap(buf))
}
```

---

### 5.3 ম্যাপ (Maps)
ম্যাপ হলো Key-Value জোড়া সংরক্ষণের জন্য হ্যাশ টেবিল।

```go
// 16_maps.go
package main

import "fmt"

func main() {
    userRoles := map[string]string{
        "goutom": "Admin",
        "rahim":  "User",
    }

    // কী (Key) চেক করার কমা-ওকে নিয়ম (comma-ok idiom)
    role, exists := userRoles["goutom"]
    fmt.Printf("রোল: %s (বিদ্যমান: %t)\n", role, exists)

    // মুছে ফেলা
    delete(userRoles, "rahim")
}
```

---

### 5.4 স্ট্রাক্ট এবং কম্পোজিশন (Structs & Composition)
Go-তে ক্লাস নেই, এর পরিবর্তে Struct এবং Composition (Embedding) ব্যবহার করা হয়।

```go
// 17_structs.go
package main

import "fmt"

type Address struct {
    City    string
    Country string
}

type User struct {
    ID      int
    Name    string
    Address // Embedded struct (Composition)
}

func main() {
    u := User{
        ID:   101,
        Name: "গৌতম রায়",
        Address: Address{
            City:    "ঢাকা",
            Country: "বাংলাদেশ",
        },
    }

    fmt.Printf("ইউজার: %+v\n", u)
    // সরাসরি এমবেডেড ফিল্ড এক্সেস (Promoted field)
    fmt.Printf("নাম: %s, শহর: %s\n", u.Name, u.City)
}
```

---

## 6. পয়েন্টার

### 6.1 পয়েন্টার বেসিক্স (`&` এবং `*`)
পয়েন্টার মেমোরি অ্যাড্রেস ধরে রাখে।
- `&`: কোনো ভ্যারিয়েবলের মেমোরি ঠিকানা নির্দেশ করে।
- `*`: ওই মেমোরি ঠিকানায় থাকা মূল মান উদ্ধার/পরিবর্তন করে।

```go
// 18_pointers.go
package main

import "fmt"

func main() {
    val := 42
    var ptr *int = &val // val-এর ঠিকানা ptr-এ রাখা হলো

    fmt.Println("মূল মান:", val)
    fmt.Println("মেমোরি অ্যাড্রেস (&val):", ptr)
    fmt.Println("পয়েন্টারের ভিতরের মান (*ptr):", *ptr)

    // পয়েন্টার দিয়ে মূল মান পরিবর্তন
    *ptr = 100
    fmt.Println("নতুন মান:", val)
}
```

---

### 6.2 পাশ-বাই-ভ্যালু বনাম পাশ-বাই-পয়েন্টার
Go সম্পূর্ণভাবে **pass-by-value**। ফাংশনের ভেতরে মূল ভ্যারিয়েবল পরিবর্তন করতে চাইলে তার পয়েন্টার পাস করতে হয়।

```go
// 19_pass_by_pointer.go
package main

import "fmt"

func updateByValue(x int) {
    x = 500 // শুধুমাত্র লোকাল কপি পরিবর্তন করবে
}

func updateByPointer(x *int) {
    *x = 500 // মূল ভ্যারিয়েবল পরিবর্তন করবে
}

func main() {
    number := 50

    updateByValue(number)
    fmt.Println("updateByValue এর পর:", number) // ৫০ থাকবে

    updateByPointer(&number)
    fmt.Println("updateByPointer এর পর:", number) // ৫০০ হয়ে যাবে
}
```

---

## 7. মেথড এবং ইন্টারফেস

### 7.1 ভ্যালু রিসিভার বনাম পয়েন্টার রিসিভার
স্ট্রাক্টের সাথে ফাংশন যুক্ত করাকেই মেথড বলে।

```go
// 20_methods.go
package main

import "fmt"

type Circle struct {
    Radius float64
}

// Value Receiver (মূল স্ট্রাক্ট পরিবর্তন করতে পারে না)
func (c Circle) Area() float64 {
    return 3.1416 * c.Radius * c.Radius
}

// Pointer Receiver (মূল স্ট্রাক্ট পরিবর্তন করতে পারে)
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

func main() {
    c := Circle{Radius: 5}
    fmt.Println("আগের ক্ষেত্রফল:", c.Area())

    c.Scale(2)
    fmt.Println("নতুন ব্যাসার্ধ:", c.Radius)
    fmt.Println("নতুন ক্ষেত্রফল:", c.Area())
}
```

---

### 7.2 ইন্টারফেস এবং ইমপ্লিসিট ইমপ্লিমেন্টেশন
Go-তে `implements` কিওয়ার্ড নেই। কোনো টাইপ যদি ইন্টারফেসের সব মেথড বাস্তবায়ন করে, তবে তা স্বয়ংক্রিয়ভাবে সেই ইন্টারফেসের অন্তর্ভুক্ত হয়।

```go
// 21_interfaces.go
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Println("ক্ষেত্রফল:", s.Area())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    printArea(rect)
}
```

---

### 7.3 টাইপ অ্যাসারশন এবং টাইপ সুইচ

```go
// 22_type_switch.go
package main

import "fmt"

func inspectType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("এটি একটি Integer:", v)
    case string:
        fmt.Println("এটি একটি String:", v)
    default:
        fmt.Printf("অন্যান্য টাইপ: %T\n", v)
    }
}

func main() {
    inspectType(100)
    inspectType("গো ল্যাঙ্গুয়েজ")
    inspectType(3.14)
}
```

---

## 8. Concurrency (কনকারেন্সি)

### 8.1 গোরুটিন (Goroutines)
গোরুটিন হলো হালকা থ্রেড যা Go রানটাইম ম্যানেজ করে। `go` কিওয়ার্ড দিয়ে শুরু করতে হয়।

```go
// 23_goroutines.go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("ব্যাকগ্রাউন্ড গোরুটিন থেকে হ্যালো!")
}

func main() {
    go sayHello() // আলাদা গোরুটিনে চলবে
    time.Sleep(100 * time.Millisecond) // মূল থ্রেডকে কিছুক্ষণ থামিয়ে রাখা
}
```

---

### 8.2 চ্যানেল (Channels)
গোরুটিনগুলোর মধ্যে ডেটা আদান-প্রদান এবং সিঙ্ক্রোনাইজেশনের মাধ্যম হলো চ্যানেল।

```go
// 24_channels.go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "কাজ শেষ!" // চ্যানেলে ডেটা পাঠানো
    }()

    msg := <-ch // চ্যানেল থেকে ডেটা গ্রহণ
    fmt.Println("বার্তা প্রাপ্তি:", msg)
}
```

---

### 8.3 `sync.WaitGroup` এবং `sync.Mutex`

```go
// 25_sync.go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
    defer wg.Done()

    c.mu.Lock()   // রেস কন্ডিশন এড়াতে লক করা
    c.value++
    c.mu.Unlock() // আনলক করা
}

func main() {
    var wg sync.WaitGroup
    c := Counter{}

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go c.Increment(&wg)
    }

    wg.Wait() // সব গোরুটিন শেষ না হওয়া পর্যন্ত অপেক্ষা
    fmt.Println("চূড়ান্ত কাউন্টার মান:", c.value)
}
```

---

## 9. মডিউল এবং স্ট্যান্ডার্ড লাইব্রেরি

### 9.1 প্যাকেজ ভিজিবিলিটি রুলস
- **Public/Exported (পাবলিক)**: নামের প্রথম অক্ষর বড় হাতের (`Capitalized`, যেমন `fmt.Println`) হলে অন্য প্যাকেজ থেকে দেখা যাবে।
- **Private/Unexported (প্রাইভেট)**: নামের প্রথম অক্ষর ছোট হাতের (`lowercase`, যেমন `myVar`) হলে কেবল একই প্যাকেজের ভিতরেই ব্যবহার করা যাবে।

### 9.2 প্রয়োজনীয় স্ট্যান্ডার্ড প্যাকেজ

```go
// 26_stdlib.go
package main

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

func main() {
    // ১. স্ট্রিং প্রসেসিং (strings)
    str := "  Go Programming  "
    fmt.Println("অতিরিক্ত স্পেস ছাঁটাই:", strings.TrimSpace(str))

    // ২. টাইপ রূপান্তর (strconv)
    num, _ := strconv.Atoi("500")
    fmt.Println("স্ট্রিং থেকে int:", num+100)

    // ৩. সময় পরিচালনা (time)
    now := time.Now()
    fmt.Println("ফরমেটেড সময়:", now.Format("2006-01-02 15:04:05"))
}
```

---

## উপসংহার (Conclusion)
অভিনন্দন! আপনি গো (Golang) প্রোগ্রামিংয়ের বেসিক গাইডলাইন সম্পূর্ণ সম্পন্ন করেছেন।
এখন আপনি মৌলিক প্রোগ্রাম, কমান্ট লাইন অ্যাপ বা সহজ REST API তৈরি করতে প্রস্তুত!
