# 🚀 Ultimate Go (Golang) Programming Course 🐹

![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.style=for-the-badge)
![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-brightgreen.svg?style=for-the-badge)
![Language](https://img.shields.io/badge/Bilingual-English%20%7C%20Bangla-blue?style=for-the-badge)

Welcome to the **Ultimate Go (Golang) Programming Basic Course Repository**! This repository is designed to take anyone from **zero to proficient** in Go programming with clear explanations, runnable code snippets, modular project architecture, and complete bilingual support (**English 🇬🇧 & Bangla 🇧🇩**).

---

## 📌 Quick Navigation (সূচিপত্র)

| # | Topic | English Guide | বাংলা গাইড | Source Code |
|---|-------|---------------|------------|-------------|
| **01** | **Introduction & Program Structure** | [English Section 1](golang_basic_course.md#1-introduction--program-structure) | [বাংলা অধ্যায় ১](golang_basic_course_bangla.md#1-সূচনা-ও-প্রোগ্রাম-স্ট্রাকচার) | [`lesson01/`](lesson01/01_introduction.go) |
| **02** | **Variables, Constants & Data Types** | [English Section 2](golang_basic_course.md#2-variables-constants--data-types) | [বাংলা অধ্যায় ২](golang_basic_course_bangla.md#2-ভ্যারিয়েবল-কনস্ট্যান্ট-এবং-ডেটা-টাইপ) | [`lesson02/`](lesson02/02_variables_types.go) |
| **03** | **Control Flow (`if`, `switch`, `for`)** | [English Section 3](golang_basic_course.md#3-control-flow) | [বাংলা অধ্যায় ৩](golang_basic_course_bangla.md#3-কন্ট্রোল-ফ্লো) | [`lesson03/`](lesson03/03_control_flow.go) |
| **04** | **Functions & Error Handling (`defer`, `panic`)** | [English Section 4](golang_basic_course.md#4-functions--error-handling) | [বাংলা অধ্যায় ৪](golang_basic_course_bangla.md#4-ফাংশন-এবং-এরর-হ্যান্ডলিং) | [`lesson04/`](lesson04/04_functions_errors.go) |
| **05** | **Data Structures (Slices, Maps, Structs)** | [English Section 5](golang_basic_course.md#5-data-structures) | [বাংলা অধ্যায় ৫](golang_basic_course_bangla.md#5-ডেটা-স্ট্রাকচার) | [`lesson05/`](lesson05/05_data_structures.go) |
| **06** | **Pointers (`&`, `*`, Pass-by-Value/Pointer)** | [English Section 6](golang_basic_course.md#6-pointers) | [বাংলা অধ্যায় ৬](golang_basic_course_bangla.md#6-পয়েন্টার) | [`lesson06/`](lesson06/06_pointers.go) |
| **07** | **Methods & Interfaces** | [English Section 7](golang_basic_course.md#7-methods--interfaces) | [বাংলা অধ্যায় ৭](golang_basic_course_bangla.md#7-মেথড-এবং-ইন্টারফেস) | [`lesson07/`](lesson07/07_methods_interfaces.go) |
| **08** | **Concurrency (Goroutines, Channels, Mutex)** | [English Section 8](golang_basic_course.md#8-concurrency) | [বাংলা অধ্যায় ৮](golang_basic_course_bangla.md#8-কনকারেন্সি) | [`lesson08/`](lesson08/08_concurrency.go) |
| **09** | **Modules & Standard Library** | [English Section 9](golang_basic_course.md#9-modules--standard-library) | [বাংলা অধ্যায় ৯](golang_basic_course_bangla.md#9-মডিউল-এবং-স্ট্যান্ডার্ড-লাইব্রেরি) | [`lesson09/`](lesson09/09_modules_stdlib.go) |
| 🌟 | **Complete Master Example** | [Full Guide](golang_basic_course.md) | [বাংলা গাইড](golang_basic_course_bangla.md) | [`bangla_example/`](bangla_example/bangla_master_example.go) |

---

## ⚡ Quick Start Guide

### 1. Prerequisites
Ensure you have **Go** installed on your system.
```bash
go version
# Output example: go version go1.26.3 windows/amd64
```

### 2. Clone the Repository
```bash
git clone https://github.com/your-username/go-basic-course.git
cd go-basic-course
```

### 3. Run Any Lesson Instantly

```bash
# Run Lesson 01 (Hello World & Program Structure)
go run ./lesson01/01_introduction.go

# Run Lesson 05 (Data Structures: Slices & Maps)
go run ./lesson05/05_data_structures.go

# Run Lesson 08 (Concurrency: Goroutines & Channels)
go run ./lesson08/08_concurrency.go

# Run the Master All-in-One Example
go run ./bangla_example/bangla_master_example.go
```

---

## 📁 Repository Architecture

```text
.
├── 📄 README.md                        # Master GitHub Repository Front Page
├── 📄 golang_basic_course.md           # Full English Documentation & Tutorial
├── 📄 golang_basic_course_bangla.md    # Full Bangla Documentation & Tutorial (সম্পূর্ণ বাংলা গাইড)
├── 📄 go.mod                           # Go Module Definition
├── 📄 .gitignore                       # Standard Go Git Ignore configuration
├── 📁 lesson01/                        # Lesson 1: Introduction & Packages
│   └── 📄 01_introduction.go
├── 📁 lesson02/                        # Lesson 2: Variables, Types & Constants
│   └── 📄 02_variables_types.go
├── 📁 lesson03/                        # Lesson 3: Control Flow (if, switch, for)
│   └── 📄 03_control_flow.go
├── 📁 lesson04/                        # Lesson 4: Functions & Errors
│   └── 📄 04_functions_errors.go
├── 📁 lesson05/                        # Lesson 5: Arrays, Slices, Maps & Structs
│   └── 📄 05_data_structures.go
├── 📁 lesson06/                        # Lesson 6: Pointers & Memory Management
│   └── 📄 06_pointers.go
├── 📁 lesson07/                        # Lesson 7: Methods & Interfaces
│   └── 📄 07_methods_interfaces.go
├── 📁 lesson08/                        # Lesson 8: Goroutines, Channels & Sync
│   └── 📄 08_concurrency.go
├── 📁 lesson09/                        # Lesson 9: Go Modules & Standard Library
│   └── 📄 09_modules_stdlib.go
└── 📁 bangla_example/                  # All-in-One Master Script
    └── 📄 bangla_master_example.go
```

---

## 📚 Key Features of Go Covered

- **Zero-Value System**: Learn how Go handles memory safely without uninitialized pointer crashes.
- **Explicit Error Handling**: Master `(value, error)` returns over standard exception throwing.
- **Composition over Inheritance**: Struct embedding for flexible OOP design.
- **CSP-style Concurrency**: High-performance multi-threading using Goroutines and Channels.
- **Clean Project Architecture**: No circular dependencies or package redeclaration errors.

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check [issues page](https://github.com/your-username/go-basic-course/issues).

---

## 📜 License

This repository is licensed under the [MIT License](LICENSE). Feel free to use, modify, and share it freely!

*Happy Coding in Go! 🐹*
