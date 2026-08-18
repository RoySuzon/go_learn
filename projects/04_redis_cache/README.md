# ⚡ Redis Caching & Cache-Aside Pattern in Go

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
**Redis** হলো একটি অতি-দ্রুত In-Memory Key-Value ডাটাবেস। রিয়েল-টাইম হাই-ট্রাফিক অ্যাপ্লিকেশনে মূল ডাটাবেসের (PostgreSQL/MySQL) ওপর চাপ কমাতে **Cache-Aside Pattern** ব্যবহার করা হয়।

---

## 🎨 Cache-Aside Pattern Diagram (ক্যাশ আর্কিটেকচার)

```text
               ┌──────────────┐
               │  Go Service  │
               └──────┬───────┘
                      │
            1. Read Cache (Get)
                      │
         ┌────────────┴────────────┐
         ▼                         ▼
   [ Cache HIT ]             [ Cache MISS ]
  (Return ~0.1ms)                  │
                         2. Query Database (100ms)
                                   │
                         3. Update Cache & Return
```

---

## 🚀 How to Run (কোড চালনার নিয়ম)

```bash
go run ./projects/04_redis_cache/main.go
```
