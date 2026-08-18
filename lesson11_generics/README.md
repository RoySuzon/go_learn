# Lesson 11: Generics in Go (Go 1.18+ টাইপ প্যারামিটার)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
Go 1.18 সংস্করণ থেকে যুক্ত হওয়া **Generics (টাইপ প্যারামিটার)** একই ফাংশন বা ডেটা স্ট্রাকচারকে বিভিন্ন ডেটা টাইপের (int, float, string) জন্য বারবার কোড না লিখে পুনঃব্যবহারযোগ্য (Reusable) করার সুবিধা দেয়।

---

## 🔑 Key Concepts (মূল ধারণাসমূহ)

1. **Type Parameters `[T any]`**: যেকোনো ডেটা টাইপ গ্রহণকারী জেনেরিক টাইপ প্যারামিটার।
2. **Constraints (`comparable`, Custom Sets)**:
   > [!NOTE]
   > `comparable` শুধু সমতা চেকে (`==`, `!=`) ব্যবহৃত হয়। কাস্টম কনস্ট্রেইন্ট যেমন `~int | ~float64` দিয়ে গাণিতিক তুলনা করা যায়।
3. **Generic Data Structures (`Stack[T]`)**: যেকোনো টাইপের জন্য সুরক্ষিত স্ট্যাক বা কিউ তৈরি।

---

## 💻 Code Example (কোড উদাহরণ)

```go
func Max[T ~int | ~float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./lesson11_generics/11_generics.go
```
