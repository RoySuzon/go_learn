# 🗺️ Ultimate Go (Golang) Learning Roadmap (step-by-step Hierarchy)

যদি আপনি গো (Go) প্রোগ্রামিং ভাষায় নতুন হন, তবে **একদম জিরো থেকে প্রফেশনাল সফটওয়্যার ইঞ্জিনিয়ার** হওয়া পর্যন্ত নিচের ধাপে ধাপে (Level by Level) শেখার সিকোয়েন্স বা রোডম্যাপ অনুসরণ করুন:

---

## 📌 রোডম্যাপ হাইরার্কি (Hierarchical Flowchart)

```text
 [Level 01: Core Syntax]
       │
       ▼  (01_introduction ──► 02_variables_types ──► 03_control_flow)
 [Level 02: Functions & Memory]
       │
       ▼  (04_functions_errors ──► 05_data_structures ──► 06_pointers)
 [Level 03: OOP & Concurrency]
       │
       ▼  (07_methods_interfaces ──► 08_concurrency ──► 09_modules_stdlib)
 [Level 04: Quality & Testing]
       │
       ▼  (10_unit_testing_test.go - Unit & Table Tests)
 [Level 05: DSA & HackerRank]
       │
       ▼  (Linked List ──► Stack/Queue ──► Trees/Graphs ──► HackerRank Problems)
 [Level 06: Real-World Projects]
       │
       ▼  (CLI Todo App ──► REST API Server with JWT & PostgreSQL)
 [Level 07: DevOps & Production]
          (Multi-Stage Docker ──► Docker Compose ──► GitHub CI/CD Actions)
```

---

## 🎯 ধাপে ধাপে শেখার নির্দেশিকা (Step-by-Step Action Plan)

### 🟢 Phase 1: Go Core Foundation (মৌলিক ভিত্তি)
> **লক্ষ্য:** গো ভাষায় সিনট্যাক্স, ভ্যারিয়েবল এবং বেসিক লজিক লেখা শেখা।

1. **Step 01: Hello World & Structure**
   - 📂 Folder: [`lesson01/`](lesson01/01_introduction.go)
   - 📖 Concept: `package main`, `import`, `func main()`, `fmt.Println`.
   - 🏃 Command: `go run ./lesson01/01_introduction.go`

2. **Step 02: Variables, Constants & Data Types**
   - 📂 Folder: [`lesson02/`](lesson02/02_variables_types.go)
   - 📖 Concept: `var`, `:=`, Zero values, `const`, `iota`, Type Casting.
   - 🏃 Command: `go run ./lesson02/02_variables_types.go`

3. **Step 03: Control Flow (`if`, `switch`, `for`)**
   - 📂 Folder: [`lesson03/`](lesson03/03_control_flow.go)
   - 📖 Concept: Conditionals with short statements, Tagless switch, Only `for` loops.
   - 🏃 Command: `go run ./lesson03/03_control_flow.go`

---

### 🟡 Phase 2: Memory & Data Management (মেমোরি ও ডেটা হ্যান্ডলিং)
> **লক্ষ্য:** অবজেক্ট হ্যান্ডলিং, স্লাইস, ম্যাপ এবং পয়েন্টার দিয়ে মেমোরি কন্ট্রোল শেখা।

4. **Step 04: Functions & Error Handling**
   - 📂 Folder: [`lesson04/`](lesson04/04_functions_errors.go)
   - 📖 Concept: Multiple returns `(result, error)`, Closures, `defer`, `panic`, `recover`.
   - 🏃 Command: `go run ./lesson04/04_functions_errors.go`

5. **Step 05: Data Structures (Slices, Maps, Structs)**
   - 📂 Folder: [`lesson05/`](lesson05/05_data_structures.go)
   - 📖 Concept: Fixed Arrays vs Dynamic Slices (`make`, `append`), Maps (`comma-ok`), Struct composition.
   - 🏃 Command: `go run ./lesson05/05_data_structures.go`

6. **Step 06: Pointers & Pass-by-Value**
   - 📂 Folder: [`lesson06/`](lesson06/06_pointers.go)
   - 📖 Concept: `&` (Address-of), `*` (Dereference), Pass-by-value vs Pass-by-pointer.
   - 🏃 Command: `go run ./lesson06/06_pointers.go`

---

### 🔵 Phase 3: Object-Oriented & Concurrency (মেথড, ইন্টারফেস ও কনকারেন্সি)
> **লক্ষ্য:** হাই-পারফরম্যান্স মাল্টি-থ্রেডেড কনকারেন্ট কোড লেখা।

7. **Step 07: Methods & Interfaces**
   - 📂 Folder: [`lesson07/`](lesson07/07_methods_interfaces.go)
   - 📖 Concept: Value vs Pointer receivers, Implicit interface implementation, Type switches.
   - 🏃 Command: `go run ./lesson07/07_methods_interfaces.go`

8. **Step 08: Concurrency (Goroutines & Channels)**
   - 📂 Folder: [`lesson08/`](lesson08/08_concurrency.go)
   - 📖 Concept: `go` routines, Unbuffered & Buffered channels, `select`, `sync.WaitGroup`, `sync.Mutex`.
   - 🏃 Command: `go run ./lesson08/08_concurrency.go`

9. **Step 09: Go Modules & Standard Library**
   - 📂 Folder: [`lesson09/`](lesson09/09_modules_stdlib.go)
   - 📖 Concept: Capitalized (Public) vs lowercase (Private) export rules, `strings`, `strconv`, `time`, `os`.
   - 🏃 Command: `go run ./lesson09/09_modules_stdlib.go`

---

### 🟠 Phase 4: Production Quality & Unit Testing (টেস্টিং ও কোয়ালিটি)
> **লক্ষ্য:** বাণিজ্যিক মানের স্বয়ংক্রিয় টেস্ট কোড লেখা।

10. **Step 10: Unit Testing & Benchmarks**
    - 📂 Folder: [`lesson10/`](lesson10/10_unit_testing_test.go)
    - 📖 Concept: `testing.T`, Table-Driven Tests, `testing.B` benchmarks.
    - 🏃 Command: `go test -v ./lesson10/...`

---

### 🟣 Phase 5: Data Structures & Algorithms (অ্যালগরিদম ও সমস্যা সমাধান)
> **লক্ষ্য:** কোডিং ইন্টারভিউ ও সফটওয়্যার ইঞ্জিনিয়ারিং প্রবলেম সলভিংয়ে দক্ষ হওয়া।

11. **Step 11: Core DSA Modules**
    - 📂 Folder: [`dsa/`](dsa/)
    - 📖 Topics: Linked List, Stack, Queue, Binary Search Tree, Quick Sort, Merge Sort.
    - 🏃 Command: `go test -v ./dsa/...`

12. **Step 12: Advanced DSA & HackerRank Challenges**
    - 📂 Folder: [`dsa/hackerrank/`](dsa/hackerrank/)
    - 📖 Challenges: Two Sum, Balanced Brackets, Reverse Linked List, Cycle Detection, Number of Islands, Coin Change.
    - 🏃 Command: `go run ./dsa/hackerrank/01_two_sum/main.go`

---

### 🔴 Phase 6: Real-World Applications & Web Projects (বাস্তবমুখী প্রজেক্টসমূহ)
> **লক্ষ্য:** পূর্ণাঙ্গ সফটওয়্যার এবং ব্যাকএন্ড সার্ভিস তৈরি করা।

13. **Step 13: CLI Todo Application**
    - 📂 Folder: [`projects/02_cli_todo/`](projects/02_cli_todo/main.go)
    - 📖 Concept: Command-line flags, JSON File Persistence.
    - 🏃 Command: `go run ./projects/02_cli_todo/main.go -add "Task"`

14. **Step 14: RESTful Web API with JWT Auth & PostgreSQL**
    - 📂 Folder: [`projects/01_rest_api/`](projects/01_rest_api/main.go)
    - 📖 Concept: `net/http` router, JWT Bearer Token Auth, Password Hashing, GORM ORM + PostgreSQL.
    - 🏃 Command: `go run ./projects/01_rest_api/main.go`

---

### 🚀 Phase 7: DevOps & Production Deployment (ক্লাউড ও ডেভঅপস)
> **লক্ষ্য:** ক্লাউড সার্ভারে অটোমেটেড কোড ডিপ্লয়মেন্ট।

15. **Step 15: Docker Containerization & Docker Compose**
    - 📄 Files: [`Dockerfile`](Dockerfile), [`docker-compose.yml`](docker-compose.yml)
    - 🏃 Command: `docker compose up -d`

16. **Step 16: GitHub Actions CI/CD Pipeline**
    - 📄 File: [`.github/workflows/go.yml`](.github/workflows/go.yml)
    - 📖 Concept: Automatic build & test verification on GitHub push.

---

🎉 **Congratulations! Following this hierarchy makes you a complete, job-ready Go Backend & Systems Engineer!**
