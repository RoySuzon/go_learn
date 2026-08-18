# Lesson 07: Methods & Interfaces (মেথড ও ইন্টারফেস)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
Go-তে `implements` কিওয়ার্ড ছাড়াই টাইপসমূহ স্বয়ংক্রিয়ভাবে (Implicitly) ইন্টারফেস বাস্তবায়ন করে।

---

## 🔑 Key Concepts (মূল ধারণাসমূহ)

1. **Value vs Pointer Receivers**:
   - Value Receiver: মূল স্ট্রাক্ট পরিবর্তন করে না।
   - Pointer Receiver: মূল স্ট্রাক্টের মান পরিবর্তন করতে পারে।
2. **Implicit Interfaces**: ইন্টারফেসের সব মেথড কোনো টাইপে ডিক্লেয়ার করা থাকলে তা স্বয়ংক্রিয়ভাবে ওই ইন্টারফেস মেনে চলে।
3. **Type Switch & Type Assertion**: `interface{}` বা `any` টাইপ থেকে মূল টাইপ পরীক্ষা করা (`val, ok := i.(string)`).

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./lesson07/07_methods_interfaces.go
```
