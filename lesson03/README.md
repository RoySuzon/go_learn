# Lesson 03: Control Flow (কন্ট্রোল ফ্লো: `if`, `switch`, `for`)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
Go-তে কন্ডিশনাল লজিক এবং একমাত্র লুপ `for` ব্যবহারের নিয়ম।

---

## 🔑 Key Concepts (মূল ধারণাসমূহ)

1. **`if` Short Initialization**: কন্ডিশন চেকের আগে ভ্যারিয়েবল ডিক্লেয়ার করা যায় (`if limit := 100; score < limit`)।
2. **Switch without `break`**: Go-তে `switch` কেসে আলাদা করে `break` লিখতে হয় না (স্বয়ংক্রিয় break)।
3. **Tagless Switch**: কন্ডিশন ছাড়া `switch` স্টেটমেন্ট `if-else` চেইনের মতো কাজ করে।
4. **Only `for` Loop**: Go-তে কোনো `while` বা `do-while` নেই; সব লুপই `for` দিয়ে লেখা হয়।

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./lesson03/03_control_flow.go
```
