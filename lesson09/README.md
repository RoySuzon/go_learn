# Lesson 09: Modules & Standard Library (মডিউল ও স্ট্যান্ডার্ড লাইব্রেরি)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
Go-তে মডিউল ব্যবস্থাপনা এবং সবচেয়ে বেশি ব্যবহৃত স্ট্যান্ডার্ড প্যাকেজসমূহ (`strings`, `strconv`, `time`, `os`)।

---

## 🔑 Key Concepts (মূল ধারণাসমূহ)

1. **Package Visibility Rules**:
   - `Capitalized`: Public / Exported (অন্য প্যাকেজ থেকে দেখা যাবে)
   - `lowercase`: Private / Unexported (একই প্যাকেজে সীমাবদ্ধ)
2. **Go Mod Commands**: `go mod init`, `go mod tidy`
3. **Standard Packages**:
   - `strings`: স্ট্রিং প্রসেসিং
   - `strconv`: টাইপ কনভার্সন
   - `time`: তারিখ ও সময় ফরমেটিং
   - `os`: অপারেটিং সিস্টেম ইন্টারফেস

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./lesson09/09_modules_stdlib.go
```
