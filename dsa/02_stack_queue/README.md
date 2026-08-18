# 02. Stack & Queue (স্ট্যাক এবং কিউ)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
- **Stack (স্ট্যাক)**: **LIFO (Last-In, First-Out)** নীতি মেনে চলে (যেমন: প্লেটের স্তূপ, Undo/Redo ফিচার)।
- **Queue (কিউ)**: **FIFO (First-In, First-Out)** নীতি মেনে চলে (যেমন: টিকিটের লাইন, রিকুয়েস্ট কিউ)।

---

## 🎨 Visualization Diagrams (চিত্রায়ন)

### Stack (LIFO):
```text
  PUSH (300)  ↓  ↑  POP ()
           │ 300 │  <- Top
           │ 200 │
           │ 100 │
           └─────┘
```

### Queue (FIFO):
```text
  ENQUEUE (30) ──► [ 30 │ 20 │ 10 ] ──► DEQUEUE (10)
                    Rear       Front
```

---

## ⚙️ Operations Summary (অপারেশনসমূহ)

### Stack Operations:
- `Push(val)`: নতুন উপাদান উপরে যুক্ত করা ($O(1)$)।
- `Pop()`: সবশেষ যুক্ত হওয়া উপাদান বের করে আনা ($O(1)$)।
- `Peek()`: উপরের উপাদানটি না সরিয়ে শুধু দেখা ($O(1)$)।

### Queue Operations:
- `Enqueue(val)`: লাইনের শেষে যুক্ত করা ($O(1)$)।
- `Dequeue()`: লাইনের প্রথম উপাদান বের করে আনা ($O(1)$)।

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)

| Operation | Time Complexity | Space Complexity |
| :--- | :--- | :--- |
| **Stack Push / Pop / Peek** | $O(1)$ | $O(N)$ |
| **Queue Enqueue / Dequeue** | $O(1)$ | $O(N)$ |

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/02_stack_queue/main.go
```
