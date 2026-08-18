# 🧠 Data Structures & Algorithms (DSA Master Module in Go)

স্বাগতম **গো (Golang) ডেটা স্ট্রাকচার এবং অ্যালগরিদম মাস্টার মডিউল**-এ। এই ফোল্ডারে সফ্টওয়্যার ইঞ্জিনিয়ারিংয়ের মৌলিক এবং অ্যাডভান্সড সকল ডেটা স্ট্রাকচার ও অ্যালগরিদম আলাদা আলাদা ফোল্ডারে কোড ও ১০০% বাংলা নির্দেশিকাসহ যুক্ত করা হয়েছে।

---

## 📌 বিষয়ভিত্তিক সূচিপত্র (Folder Index)

| # | Topic | Description | Bangla Guide | Source Code | Command to Run |
|---|-------|-------------|--------------|-------------|----------------|
| **01** | **Linked List** | Singly Linked List (Head, Tail, Delete, Search) | [বাংলা গাইড](01_linked_list/README.md) | [`01_linked_list/main.go`](01_linked_list/main.go) | `go run ./dsa/01_linked_list/main.go` |
| **02** | **Stack & Queue** | LIFO Stack & FIFO Queue Operations | [বাংলা গাইড](02_stack_queue/README.md) | [`02_stack_queue/main.go`](02_stack_queue/main.go) | `go run ./dsa/02_stack_queue/main.go` |
| **03** | **Binary Tree** | Binary Search Tree (BST) & In-Order Traversal | [বাংলা গাইড](03_binary_tree/README.md) | [`03_binary_tree/main.go`](03_binary_tree/main.go) | `go run ./dsa/03_binary_tree/main.go` |
| **04** | **Searching & Sorting** | Linear/Binary Search, Quick/Merge Sort | [বাংলা গাইড](04_searching_sorting/README.md) | [`04_searching_sorting/main.go`](04_searching_sorting/main.go) | `go run ./dsa/04_searching_sorting/main.go` |
| **05** | **Advanced DSA** | Trie, Min-Heap, Graph BFS, Sliding Window | [বাংলা গাইড](05_advanced_dsa/README.md) | [`05_advanced_dsa/main.go`](05_advanced_dsa/main.go) | `go run ./dsa/05_advanced_dsa/main.go` |
| 🏆 | **HackerRank Solutions** | 6 Standalone HackerRank Challenge Solutions | [বাংলা প্রবলেমস গাইড](hackerrank/README.md) | [`hackerrank/`](hackerrank/) | `go run ./dsa/hackerrank/01_two_sum/main.go` |

---

## 📊 Big-O Complexity Quick Reference

| Structure / Algorithm | Search | Insertion | Deletion | Worst Time | Space |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Linked List** | $O(N)$ | $O(1)$ | $O(1)$ | $O(N)$ | $O(N)$ |
| **Stack / Queue** | $O(N)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(N)$ |
| **Binary Search Tree** | $O(\log N)$ | $O(\log N)$ | $O(\log N)$ | $O(N)$ | $O(N)$ |
| **Trie (Prefix Tree)** | $O(L)$ | $O(L)$ | $O(L)$ | $O(L)$ | $O(N \cdot L)$ |
| **Quick / Merge Sort** | - | - | - | $O(N \log N)$ | $O(N)$ |

---

## 🧪 Run Unit Tests
```bash
go test -v ./...
```
