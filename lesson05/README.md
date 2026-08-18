# Lesson 05: Data Structures (অ্যারে, স্লাইস, ম্যাপ ও স্ট্রাক্ট)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
Go-এর বিল্ডিং-ব্লক ডেটা স্ট্রাকচারসমূহ এবং Struct Embedding দিয়ে Composition রূপায়ন।

---

## 🔑 Key Concepts (মূল ধারণাসমূহ)

1. **Arrays vs Slices**:
   - Array: নির্দিষ্ট আকার (`[3]int`)
   - Slice: গতিশীল আকার (`[]int`), `make()`, `append()` দিয়ে বাড়ানো যায়।
2. **Maps (Key-Value Dictionaries)**:
   - Comma-ok idiom দিয়ে কী (Key) অস্তিত্ব চেক করা (`val, exists := myMap[key]`)।
3. **Struct Composition (Embedding)**:
   - Go-তে ক্লাসের উত্তরাধিকার (Inheritance) নেই; তার বদলে Struct Embedding ব্যবহৃত হয়।

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./lesson05/05_data_structures.go
```
