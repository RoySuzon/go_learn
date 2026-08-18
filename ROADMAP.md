# 🗺️ Ultimate Go (Golang) Learning Roadmap (step-by-step Hierarchy)

যদি আপনি গো (Go) প্রোগ্রামিং ভাষায় নতুন হন, তবে **একদম জিরো থেকে প্রফেশনাল সিনিয়র সফটওয়্যার ইঞ্জিনিয়ার** হওয়া পর্যন্ত নিচের ধাপে ধাপে (Level by Level) শেখার সিকোয়েন্স বা রোডম্যাপ অনুসরণ করুন:

---

## 📌 রোডম্যাপ হাইরার্কি (Hierarchical Flowchart)

```text
 [Level 01: Core Syntax]
       │
       ▼  (01_introduction ──► 02_variables_types ──► 03_control_flow)
 [Level 02: Memory & Functions]
       │
       ▼  (04_functions_errors ──► 05_data_structures ──► 06_pointers)
 [Level 03: OOP & Concurrency]
       │
       ▼  (07_methods_interfaces ──► 08_concurrency ──► 09_modules_stdlib)
 [Level 04: Modern Go, Generics & Testing]
       │
       ▼  (10_unit_testing ──► 11_generics ──► 12_context)
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
1. **Step 01: Hello World & Structure** — [`lesson01/`](lesson01/README.md) (`go run ./lesson01/01_introduction.go`)
2. **Step 02: Variables & Data Types** — [`lesson02/`](lesson02/README.md) (`go run ./lesson02/02_variables_types.go`)
3. **Step 03: Control Flow (`if`, `switch`, `for`)** — [`lesson03/`](lesson03/README.md) (`go run ./lesson03/03_control_flow.go`)

### 🟡 Phase 2: Memory & Data Management (মেমোরি ও ডেটা হ্যান্ডলিং)
4. **Step 04: Functions & Error Handling** — [`lesson04/`](lesson04/README.md) (`go run ./lesson04/04_functions_errors.go`)
5. **Step 05: Data Structures (Slices, Maps, Structs)** — [`lesson05/`](lesson05/README.md) (`go run ./lesson05/05_data_structures.go`)
6. **Step 06: Pointers & Pass-by-Value** — [`lesson06/`](lesson06/README.md) (`go run ./lesson06/06_pointers.go`)

### 🔵 Phase 3: Object-Oriented & Concurrency (মেথড, ইন্টারফেস ও কনকারেন্সি)
7. **Step 07: Methods & Interfaces** — [`lesson07/`](lesson07/README.md) (`go run ./lesson07/07_methods_interfaces.go`)
8. **Step 08: Concurrency (Goroutines & Channels)** — [`lesson08/`](lesson08/README.md) (`go run ./lesson08/08_concurrency.go`)
9. **Step 09: Go Modules & Standard Library** — [`lesson09/`](lesson09/README.md) (`go run ./lesson09/09_modules_stdlib.go`)

### 🟠 Phase 4: Quality, Generics & Context (টেস্টিং, জেনেরিক্স ও কনটেক্সট)
10. **Step 10: Unit Testing & Benchmarks** — [`lesson10/`](lesson10/README.md) (`go test -v ./lesson10/...`)
11. **Step 11: Generics in Go (Go 1.18+)** — [`lesson11_generics/`](lesson11_generics/README.md) (`go run ./lesson11_generics/11_generics.go`)
12. **Step 12: Context Package & Timeout Cancellation** — [`lesson12_context/`](lesson12_context/README.md) (`go run ./lesson12_context/12_context.go`)

### 🟣 Phase 5: Data Structures & Algorithms (অ্যালগরিদম ও প্রবলেম সলভিং)
13. **Step 13: Core DSA Modules** — [`dsa/`](dsa/README.md) (`go test -v ./dsa/...`)
14. **Step 14: HackerRank 6 Top Problem Patterns** — [`dsa/hackerrank/`](dsa/hackerrank/README.md)

### 🔴 Phase 6: Real-World Systems Engineering Projects (৮টি ব্যাকএন্ড প্রজেক্ট)
15. **Step 15: RESTful Web API (JWT Auth + GORM PostgreSQL)** — [`projects/01_rest_api/`](projects/01_rest_api/README.md)
16. **Step 16: CLI Todo Manager App** — [`projects/02_cli_todo/`](projects/02_cli_todo/README.md)
17. **Step 17: High-Speed gRPC Microservice** — [`projects/03_grpc_service/`](projects/03_grpc_service/README.md)
18. **Step 18: Distributed Redis Caching** — [`projects/04_redis_cache/`](projects/04_redis_cache/README.md)
19. **Step 19: Event Queue Broker & Worker Pool** — [`projects/05_event_queue/`](projects/05_event_queue/README.md)
20. **Step 20: Advanced SQL Database Engineering** — [`projects/06_database_advanced/`](projects/06_database_advanced/README.md)
21. **Step 21: System Design Patterns (Rate Limiter & Circuit Breaker)** — [`projects/07_system_design/`](projects/07_system_design/README.md)
22. **Step 22: Observability & Metrics (`slog` & Prometheus)** — [`projects/08_observability/`](projects/08_observability/README.md)

### 🚀 Phase 7: DevOps & Production Deployment (ক্লাউড ও ডেভঅপস)
23. **Step 23: Docker, Compose & GitHub Actions CI/CD** — [`Dockerfile`](Dockerfile), [`docker-compose.yml`](docker-compose.yml), [`.github/workflows/go.yml`](.github/workflows/go.yml)
24. **Step 24: Senior Software Engineering Architecture Guide** — [`docs/system_architecture_bangla.md`](docs/system_architecture_bangla.md)
