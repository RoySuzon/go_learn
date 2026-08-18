# 🗺️ Ultimate Go (Golang) Learning Roadmap (step-by-step Hierarchy)

যদি আপনি গো (Go) প্রোগ্রামিং ভাষায় নতুন হন, তবে **একদম জিরো থেকে প্রফেশনাল সিনিয়র সফটওয়্যার ইঞ্জিনিয়ার** হওয়া পর্যন্ত নিচের ধাপে ধাপে (Level by Level) শেখার সিকোয়েন্স বা রোডম্যাপ অনুসরণ করুন:

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
 [Level 05: Data Structures & Algorithms]
       │
       ▼  (01_linked_list ──► 02_stack_queue ──► 03_binary_tree ──► 04_searching_sorting ──► 05_advanced_dsa ──► hackerrank)
 [Level 06: Real-World Backend & Systems Engineering Projects]
       │
       ▼  (01_rest_api ──► 02_cli_todo ──► 03_grpc_service ──► 04_redis_cache ──► 05_event_queue ──► 06_database_advanced ──► 07_system_design ──► 08_observability)
 [Level 07: DevOps & Production Architecture]
          (Dockerfile ──► docker-compose.yml ──► .github/workflows/go.yml ──► system_architecture_bangla.md)
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
    - 📂 Subfolders: [`dsa/01_linked_list/`](dsa/01_linked_list/README.md), [`dsa/02_stack_queue/`](dsa/02_stack_queue/README.md), [`dsa/03_binary_tree/`](dsa/03_binary_tree/README.md), [`dsa/04_searching_sorting/`](dsa/04_searching_sorting/README.md), [`dsa/05_advanced_dsa/`](dsa/05_advanced_dsa/README.md).
    - 📖 Topics: Linked List, Stack, Queue, Binary Search Tree, Quick Sort, Merge Sort, Trie, Min-Heap, Graph BFS, Sliding Window.
    - 🏃 Command: `go test -v ./dsa/...`

12. **Step 12: HackerRank & LeetCode Standalone Challenges**
    - 📂 Folder: [`dsa/hackerrank/`](dsa/hackerrank/README.md)
    - 📖 Standalone Problems: Two Sum, Balanced Brackets, Reverse Linked List, Cycle Detection, Number of Islands, Coin Change.
    - 🏃 Command: `go run ./dsa/hackerrank/01_two_sum/main.go`

---

### 🔴 Phase 6: Real-World Systems Engineering Projects (বাস্তবমুখী প্রজেক্টসমূহ)
> **লক্ষ্য:** পূর্ণাঙ্গ মাইক্রোসার্ভিস, ডিস্ট্রিবিউটেড সিউম এবং ব্যাকএন্ড ইঞ্জিনিয়ারিং আয়ত্ত করা।

13. **Step 13: RESTful Web API (JWT Auth + GORM PostgreSQL)**
    - 📂 Folder: [`projects/01_rest_api/`](projects/01_rest_api/README.md)
    - 📖 Concept: JWT Bearer Auth, Password Hashing, GORM ORM PostgreSQL Persistence.
    - 🏃 Command: `go run ./projects/01_rest_api/main.go`

14. **Step 14: CLI Todo Application**
    - 📂 Folder: [`projects/02_cli_todo/`](projects/02_cli_todo/README.md)
    - 📖 Concept: Command-line flags, JSON File Persistence.
    - 🏃 Command: `go run ./projects/02_cli_todo/main.go -list`

15. **Step 15: High-Speed gRPC Microservice**
    - 📂 Folder: [`projects/03_grpc_service/`](projects/03_grpc_service/README.md)
    - 📖 Concept: Protocol Buffers (`proto3`), Unary RPC over HTTP/2.
    - 🏃 Command: `go run ./projects/03_grpc_service/server/main.go`

16. **Step 16: Distributed Redis Caching**
    - 📂 Folder: [`projects/04_redis_cache/`](projects/04_redis_cache/README.md)
    - 📖 Concept: Cache-Aside Pattern, TTL Expiration, Sub-millisecond hits.
    - 🏃 Command: `go run ./projects/04_redis_cache/main.go`

17. **Step 17: Event Queue Broker & Worker Pool**
    - 📂 Folder: [`projects/05_event_queue/`](projects/05_event_queue/README.md)
    - 📖 Concept: Async Message Queue, Consumer Worker Pool, Dead-Letter Queue (DLQ).
    - 🏃 Command: `go run ./projects/05_event_queue/main.go`

18. **Step 18: Advanced SQL Database Engineering**
    - 📂 Folder: [`projects/06_database_advanced/`](projects/06_database_advanced/README.md)
    - 📖 Concept: Connection Pooling (`SetMaxOpenConns`), ACID Database Transactions (`BEGIN`, `COMMIT`, `ROLLBACK`).
    - 🏃 Command: `go run ./projects/06_database_advanced/main.go`

19. **Step 19: System Design Patterns (Rate Limiter & Circuit Breaker)**
    - 📂 Folder: [`projects/07_system_design/`](projects/07_system_design/README.md)
    - 📖 Concept: Token Bucket Rate Limiter ($O(1)$) and Circuit Breaker State Machine.
    - 🏃 Command: `go run ./projects/07_system_design/main.go`

20. **Step 20: Observability, Metrics & Structured Logging**
    - 📂 Folder: [`projects/08_observability/`](projects/08_observability/README.md)
    - 📖 Concept: Structured JSON Logging (`log/slog`) and Prometheus Metrics Counter/Histogram.
    - 🏃 Command: `go run ./projects/08_observability/main.go`

---

### 🚀 Phase 7: DevOps, Containerization & Architecture (ক্লাউড ও ডেভঅপস)
> **লক্ষ্য:** ক্লাউড ইনফ্রাস্ট্রাকচার ও অটোমেটেড সিআই/সিডিPipeline।

21. **Step 21: Docker Containerization & Docker Compose**
    - 📄 Files: [`Dockerfile`](Dockerfile), [`docker-compose.yml`](docker-compose.yml)
    - 🏃 Command: `docker compose up -d`

22. **Step 22: GitHub Actions CI/CD Pipeline**
    - 📄 File: [`.github/workflows/go.yml`](.github/workflows/go.yml)
    - 📖 Concept: Automatic build & test verification on GitHub push.

23. **Step 23: Senior Software Engineering Architecture Guide**
    - 📄 File: [`docs/system_architecture_bangla.md`](docs/system_architecture_bangla.md)

---

🎉 **Congratulations! Following this 23-step hierarchy makes you a complete, job-ready Senior Go Backend & Systems Engineer!**
