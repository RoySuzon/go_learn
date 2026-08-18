# 🚀 Ultimate Go (Golang) Master Course, DSA & Architecture 🐹

[![Go CI](https://github.com/RoySuzon/go_learn/actions/workflows/go.yml/badge.svg)](https://github.com/RoySuzon/go_learn/actions/workflows/go.yml)
![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)
![Language](https://img.shields.io/badge/Bilingual-English%20%7C%20Bangla-blue?style=for-the-badge)

Welcome to the **Ultimate Go (Golang) Programming Master Course Repository**! This repository takes you from **zero to production-ready Senior Software Engineer** in Go programming with clear step-by-step hierarchy roadmap, runnable code snippets, unit tests, real-world projects, gRPC microservices, Redis caching, Message queues, System design patterns, Observability, Data Structures & Algorithms (DSA), HackerRank/LeetCode solutions, System Architecture guides, and complete bilingual support (**English 🇬🇧 & Bangla 🇧🇩**).

---

## 🗺️ Step-by-Step Learning Hierarchy Roadmap

```text
 [Phase 1: Core Syntax] ──► [Phase 2: Memory & Functions] ──► [Phase 3: OOP & Concurrency]
                                                                        │
 ┌──────────────────────────────────────────────────────────────────────┘
 ▼
 [Phase 4: Unit Testing & Quality] ──► [Phase 5: DSA & HackerRank]
                                                    │
 ┌──────────────────────────────────────────────────┘
 ▼
 [Phase 6: Real-World Backend & Systems Engineering]
  ├── REST API Server (JWT + PostgreSQL GORM)
  ├── gRPC Microservice Service & Client
  ├── Distributed Redis Caching (Cache-Aside)
  ├── Event-Driven Queue & Worker Pool (DLQ)
  ├── Advanced Database Engineering (Connection Pooling & ACID Tx)
  ├── System Design Patterns (Token Bucket Rate Limiter & Circuit Breaker)
  └── Observability & Structured Logging (log/slog & Prometheus Metrics)
       │
       ▼
 [Phase 7: DevOps & Production Deployment] (Docker ──► Compose ──► GitHub CI/CD)
```

👉 **[Read the Full Step-by-Step Hierarchy Roadmap (ROADMAP.md)](ROADMAP.md)**  
👉 **[Read System Architecture & Senior Software Engineering Guide (Bangla)](docs/system_architecture_bangla.md)**

---

## 📌 Quick Navigation (সূচিপত্র)

| # | Topic | English Guide | বাংলা গাইড | Source Code | Command to Run |
|---|-------|---------------|------------|-------------|----------------|
| **01** | **Introduction & Program Structure** | [English 1](golang_basic_course.md#1-introduction--program-structure) | [বাংলা ১](golang_basic_course_bangla.md#1-সূচনা-ও-প্রোগ্রাম-স্ট্রাকচার) | [`lesson01/`](lesson01/01_introduction.go) | `go run ./lesson01/01_introduction.go` |
| **02** | **Variables & Data Types** | [English 2](golang_basic_course.md#2-variables-constants--data-types) | [বাংলা ২](golang_basic_course_bangla.md#2-ভ্যারিয়েবল-কনস্ট্যান্ট-এবং-ডেটা-টাইপ) | [`lesson02/`](lesson02/02_variables_types.go) | `go run ./lesson02/02_variables_types.go` |
| **03** | **Control Flow (`if`, `switch`, `for`)** | [English 3](golang_basic_course.md#3-control-flow) | [বাংলা ৩](golang_basic_course_bangla.md#3-কন্ট্রোল-ফ্লো) | [`lesson03/`](lesson03/03_control_flow.go) | `go run ./lesson03/03_control_flow.go` |
| **04** | **Functions & Error Handling** | [English 4](golang_basic_course.md#4-functions--error-handling) | [বাংলা ৪](golang_basic_course_bangla.md#4-ফাংশন-এবং-এরর-হ্যান্ডলিং) | [`lesson04/`](lesson04/04_functions_errors.go) | `go run ./lesson04/04_functions_errors.go` |
| **05** | **Data Structures (Slices, Maps, Structs)** | [English 5](golang_basic_course.md#5-data-structures) | [বাংলা ৫](golang_basic_course_bangla.md#5-ডেটা-স্ট্রাকচার) | [`lesson05/`](lesson05/05_data_structures.go) | `go run ./lesson05/05_data_structures.go` |
| **06** | **Pointers & Memory Management** | [English 6](golang_basic_course.md#6-pointers) | [বাংলা ৬](golang_basic_course_bangla.md#6-পয়েন্টার) | [`lesson06/`](lesson06/06_pointers.go) | `go run ./lesson06/06_pointers.go` |
| **07** | **Methods & Interfaces** | [English 7](golang_basic_course.md#7-methods--interfaces) | [বাংলা ৭](golang_basic_course_bangla.md#7-মেথড-এবং-ইন্টারফেস) | [`lesson07/`](lesson07/07_methods_interfaces.go) | `go run ./lesson07/07_methods_interfaces.go` |
| **08** | **Concurrency (Goroutines, Channels, Mutex)** | [English 8](golang_basic_course.md#8-concurrency) | [বাংলা ৮](golang_basic_course_bangla.md#8-কনকারেন্সি) | [`lesson08/`](lesson08/08_concurrency.go) | `go run ./lesson08/08_concurrency.go` |
| **09** | **Modules & Standard Library** | [English 9](golang_basic_course.md#9-modules--standard-library) | [বাংলা ৯](golang_basic_course_bangla.md#9-মডিউল-এবং-স্ট্যান্ডার্ড-লাইব্রেরি) | [`lesson09/`](lesson09/09_modules_stdlib.go) | `go run ./lesson09/09_modules_stdlib.go` |
| **10** | **Unit Testing & Benchmarks** | [Testing Guide](golang_basic_course.md) | [টেস্টিং গাইড](golang_basic_course_bangla.md) | [`lesson10/`](lesson10/10_unit_testing_test.go) | `go test -v ./lesson10/...` |
| 🧠 | **Data Structures & Algorithms (DSA)** | [DSA Guide](dsa/README.md) | [বাংলা DSA গাইড](dsa/README.md) | [`dsa/`](dsa/) | `go test -v ./dsa/...` |
| 🏆 | **HackerRank & LeetCode Solutions** | [Solutions Guide](dsa/hackerrank/README.md) | [বাংলা প্রবলেমস গাইড](dsa/hackerrank/README.md) | [`dsa/hackerrank/`](dsa/hackerrank/) | `go run ./dsa/hackerrank/01_two_sum/main.go` |
| 🌐 | **Project 01: REST API (JWT + Postgres)** | - | [REST API গাইড](golang_basic_course_bangla.md) | [`projects/01_rest_api/`](projects/01_rest_api/main.go) | `go run ./projects/01_rest_api/main.go` |
| 📝 | **Project 02: CLI Todo Manager App** | - | [CLI App গাইড](golang_basic_course_bangla.md) | [`projects/02_cli_todo/`](projects/02_cli_todo/main.go) | `go run ./projects/02_cli_todo/main.go -list` |
| ⚡ | **Project 03: High-Speed gRPC Microservice** | - | [gRPC গাইড](projects/03_grpc_service/README.md) | [`projects/03_grpc_service/`](projects/03_grpc_service/) | `go run ./projects/03_grpc_service/server/main.go` |
| 🏎️ | **Project 04: Redis Caching (Cache-Aside)** | - | [Redis গাইড](projects/04_redis_cache/README.md) | [`projects/04_redis_cache/`](projects/04_redis_cache/main.go) | `go run ./projects/04_redis_cache/main.go` |
| 📬 | **Project 05: Event Queue & DLQ Broker** | - | [Event Queue গাইড](projects/05_event_queue/README.md) | [`projects/05_event_queue/`](projects/05_event_queue/main.go) | `go run ./projects/05_event_queue/main.go` |
| 🗄️ | **Project 06: Advanced SQL & Connection Pool** | - | [Advanced SQL গাইড](projects/06_database_advanced/README.md) | [`projects/06_database_advanced/`](projects/06_database_advanced/main.go) | `go run ./projects/06_database_advanced/main.go` |
| 🛡️ | **Project 07: System Design (Rate Limiter & CB)** | - | [System Design গাইড](projects/07_system_design/README.md) | [`projects/07_system_design/`](projects/07_system_design/main.go) | `go run ./projects/07_system_design/main.go` |
| 📊 | **Project 08: Observability (slog & Prometheus)** | - | [Observability গাইড](projects/08_observability/README.md) | [`projects/08_observability/`](projects/08_observability/main.go) | `go run ./projects/08_observability/main.go` |

---

## ⚡ Quick Start & Running Commands

```bash
# 1. Run All Unit Tests Across Entire Workspace
go test -v ./...

# 2. Run High-Speed Redis Caching Demo
go run ./projects/04_redis_cache/main.go

# 3. Run Event Queue Broker Demo
go run ./projects/05_event_queue/main.go

# 4. Run System Design Patterns (Rate Limiter & Circuit Breaker)
go run ./projects/07_system_design/main.go

# 5. Run Observability (Structured slog & Prometheus)
go run ./projects/08_observability/main.go

# 6. Docker Deployment (Containerized API & PostgreSQL Service)
docker compose up -d
```

---

## 📜 License

This repository is licensed under the [MIT License](LICENSE). Feel free to star ⭐️, fork, and learn!
