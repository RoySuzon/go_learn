# Lesson 08: Concurrency in Go (কনকারেন্সি ও গোরুটিন)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
Go-এর অন্যতম প্রধান ফিচার হলো লাইটওয়েট থ্রেড **Goroutine** এবং কমিউনিকেশনের জন্য **Channel**।

---

## 🔑 Key Concepts (মূল ধারণাসমূহ)

1. **Goroutine (`go`)**: লাইটওয়েট থ্রেড যা খুব কম মেমোরি (~2 KB) ব্যবহার করে।
2. **Channels (`make(chan T)`)**: গোরুটিনগুলোর মধ্যে নিরাপদে ডেটা আদান-প্রদানের মাধ্যম (Unbuffered vs Buffered)।
3. **`select` Statement**: একাধিক চ্যানেলের ইনপুট/আউটপুট হ্যান্ডেল করা।
4. **`sync.WaitGroup` & `sync.Mutex`**: গোরুটিন শেষ হওয়া পর্যন্ত অপেক্ষা এবং ডাটা রেস এড়াতে রেসোর্স লক করা।

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./lesson08/08_concurrency.go
```
