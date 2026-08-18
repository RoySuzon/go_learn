# 🚀 Ultimate Go (Golang) Master Course & Projects 🐹

[![Go CI](https://github.com/RoySuzon/go_learn/actions/workflows/go.yml/badge.svg)](https://github.com/RoySuzon/go_learn/actions/workflows/go.yml)
![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)
![Language](https://img.shields.io/badge/Bilingual-English%20%7C%20Bangla-blue?style=for-the-badge)

Welcome to the **Ultimate Go (Golang) Programming Master Course Repository**! This repository takes you from **zero to production-ready developer** in Go programming with clear explanations, runnable code snippets, unit tests, real-world projects, and complete bilingual support (**English 🇬🇧 & Bangla 🇧🇩**).

---

## 📌 Quick Navigation (সূচিপত্র)

| # | Topic | English Guide | বাংলা গাইড | Source Code |
|---|-------|---------------|------------|-------------|
| **01** | **Introduction & Program Structure** | [English Section 1](golang_basic_course.md#1-introduction--program-structure) | [বাংলা অধ্যায় ১](golang_basic_course_bangla.md#1-সূচনা-ও-প্রোগ্রাম-স্ট্রাকচার) | [`lesson01/`](lesson01/01_introduction.go) |
| **02** | **Variables, Constants & Data Types** | [English Section 2](golang_basic_course.md#2-variables-constants--data-types) | [বাংলা অধ্যায় ২](golang_basic_course_bangla.md#2-ভ্যারিয়েবল-কনস্ট্যান্ট-এবং-ডেটা-টাইপ) | [`lesson02/`](lesson02/02_variables_types.go) |
| **03** | **Control Flow (`if`, `switch`, `for`)** | [English Section 3](golang_basic_course.md#3-control-flow) | [বাংলা অধ্যায় ৩](golang_basic_course_bangla.md#3-কন্ট্রোল-ফ্লো) | [`lesson03/`](lesson03/03_control_flow.go) |
| **04** | **Functions & Error Handling** | [English Section 4](golang_basic_course.md#4-functions--error-handling) | [বাংলা অধ্যায় ৪](golang_basic_course_bangla.md#4-ফাংশন-এবং-এরর-হ্যান্ডলিং) | [`lesson04/`](lesson04/04_functions_errors.go) |
| **05** | **Data Structures (Slices, Maps, Structs)** | [English Section 5](golang_basic_course.md#5-data-structures) | [বাংলা অধ্যায় ৫](golang_basic_course_bangla.md#5-ডেটা-স্ট্রাকচার) | [`lesson05/`](lesson05/05_data_structures.go) |
| **06** | **Pointers (`&`, `*`, Pass-by-Value/Pointer)** | [English Section 6](golang_basic_course.md#6-pointers) | [বাংলা অধ্যায় ৬](golang_basic_course_bangla.md#6-পয়েন্টার) | [`lesson06/`](lesson06/06_pointers.go) |
| **07** | **Methods & Interfaces** | [English Section 7](golang_basic_course.md#7-methods--interfaces) | [বাংলা অধ্যায় ৭](golang_basic_course_bangla.md#7-মেথড-এবং-ইন্টারফেস) | [`lesson07/`](lesson07/07_methods_interfaces.go) |
| **08** | **Concurrency (Goroutines, Channels, Mutex)** | [English Section 8](golang_basic_course.md#8-concurrency) | [বাংলা অধ্যায় ৮](golang_basic_course_bangla.md#8-কনকারেন্সি) | [`lesson08/`](lesson08/08_concurrency.go) |
| **09** | **Modules & Standard Library** | [English Section 9](golang_basic_course.md#9-modules--standard-library) | [বাংলা অধ্যায় ৯](golang_basic_course_bangla.md#9-মডিউল-এবং-স্ট্যান্ডার্ড-লাইব্রেরি) | [`lesson09/`](lesson09/09_modules_stdlib.go) |
| **10** | **Unit Testing & Benchmarks** | [Testing Guide](golang_basic_course.md) | [টেস্টিং গাইড](golang_basic_course_bangla.md) | [`lesson10/`](lesson10/10_unit_testing_test.go) |
| 🌐 | **Project 1: RESTful Web API Server** | - | [REST API গাইড](golang_basic_course_bangla.md) | [`projects/01_rest_api/`](projects/01_rest_api/main.go) |
| 📝 | **Project 2: CLI Todo Manager App** | - | [CLI App গাইড](golang_basic_course_bangla.md) | [`projects/02_cli_todo/`](projects/02_cli_todo/main.go) |
| 🌟 | **Master All-in-One Example** | [Full Guide](golang_basic_course.md) | [বাংলা গাইড](golang_basic_course_bangla.md) | [`bangla_example/`](bangla_example/bangla_master_example.go) |

---

## ⚡ Quick Start & Running Commands

```bash
# 1. Run Unit Tests & Benchmarks
go test -v ./...
go test -bench=. ./lesson10/...

# 2. Run REST API Server
go run ./projects/01_rest_api/main.go
# Open browser / Postman: http://localhost:8080/api/books

# 3. Run CLI Todo App
go run ./projects/02_cli_todo/main.go -add "Go প্রোগ্রামিং শেখা"
go run ./projects/02_cli_todo/main.go -list
go run ./projects/02_cli_todo/main.go -complete 1

# 4. Run Bangla Master Example
go run ./bangla_example/bangla_master_example.go
```

---

## 📁 Repository Architecture

```text
.
├── 📄 README.md                        # Master GitHub Homepage
├── 📄 golang_basic_course.md           # English Course Manual
├── 📄 golang_basic_course_bangla.md    # বাংলা কোর্স গাইড
├── 📄 go.mod                           # Go Module Definition
├── 📁 .github/workflows/               # GitHub Actions CI Workflow
│   └── 📄 go.yml
├── 📁 lesson01/ to lesson09/           # Foundational Core Lessons
├── 📁 lesson10/                        # Unit Testing & Benchmarking
│   └── 📄 10_unit_testing_test.go
├── 📁 projects/
│   ├── 📁 01_rest_api/                 # RESTful HTTP API Server
│   │   └── 📄 main.go
│   └── 📁 02_cli_todo/                 # CLI Task Manager with Persistence
│       └── 📄 main.go
└── 📁 bangla_example/                  # Bangla Master Example
    └── 📄 bangla_master_example.go
```

---

## 📜 License

This repository is licensed under the [MIT License](LICENSE). Feel free to star ⭐️, fork, and learn!
