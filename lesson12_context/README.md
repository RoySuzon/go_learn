# Lesson 12: Go Context Package (কনটেক্সট ও টাইমআউট হ্যান্ডলিং)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
**`context` প্যাকেজ** গো ব্যাকএন্ড অ্যাপ্লিকেশনে টাইমআউট (Timeout), রিকোয়েস্ট বাতিল করা (Cancellation Signal), এবং থ্রেডগুলোর মধ্যে ট্র্যাকিং ভ্যালু (Trace ID / Request ID) পাস করার প্রধান হাতিয়ার।

---

## 🔑 Key Concepts (মূল ধারণাসমূহ)

1. **`context.WithTimeout(parent, duration)`**: নির্দিষ্ট সময় অতিক্রম করলে কাজ বাতিল করা (`context.DeadlineExceeded`)।
2. **`context.WithCancel(parent)`**: হাত দিয়ে ক্যানসেল সিগন্যাল পাঠানো।
3. **`context.WithValue(parent, key, val)`**: রিকোয়েস্ট আইডেন্টিফায়ার পাঠাতে।

> [!IMPORTANT]
> প্রোডাকশন ব্যাকএন্ড সার্ভিস যেমন HTTP হ্যান্ডলার, gRPC রিকোয়েস্ট, বা SQL কোয়েরি তৈরি করার সময় প্রথম আর্গুমেন্ট হিসেবে সর্বদা `ctx context.Context` পাঠানো গো-এর একটি মৌলিক নিয়ম।

---

## 💻 Code Example (কোড উদাহরণ)

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
```

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./lesson12_context/12_context.go
```
