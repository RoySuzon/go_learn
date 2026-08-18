# 🛡️ System Design Patterns: Rate Limiting & Circuit Breaker

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
অতিরিক্ত এপিআই রিকোয়েস্টের ফলে সার্ভার ক্র্যাশ হওয়া ঠেকাতে **Rate Limiting (Token Bucket)** এবং ডাউনস্ট্রিম সার্ভিস ডাউনে সম্পূর্ণ সিস্টেম ব্যর্থ হওয়া ঠেকাতে **Circuit Breaker Pattern** ব্যবহৃত হয়।

---

## 🎨 System Design Diagrams

### 1. Token Bucket Rate Limiter:
```text
           [ Refill Tokens ]
                  │
                  ▼
          ┌───────────────┐
          │ Token Bucket  │  (Capacity: 3)
          └───────┬───────┘
                  │
          Take 1 Token/Request
         /                 \
    (Tokens > 0)       (Tokens == 0)
         │                   │
    ✅ ALLOWED           ❌ BLOCKED (HTTP 429)
```

### 2. Circuit Breaker State Machine:
```text
   [ CLOSED ] ─── High Errors ───► [ OPEN ]
       ▲                              │
       │                            Timeout
       └────── Success ──────── [ HALF-OPEN ]
```

---

## 🚀 How to Run (কোড চালনার নিয়ম)

```bash
go run ./projects/07_system_design/main.go
```
