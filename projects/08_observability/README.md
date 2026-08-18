# 📊 Observability, Structured Logging (`slog`) & Prometheus Metrics

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
প্রোডাকশন সার্ভারে সিস্টেমের স্বাস্থ্যের অবস্থা, রেসপন্স টাইম (Latency) এবং ত্রুটি মনিটর করার প্রক্রিয়াকে **Observability** বলা হয়।

---

## ⚙️ Key Concepts Covered

1. **Structured JSON Logging (`log/slog`)**:
   - ডেটাকুটির স্ট্যাকড্রাইভার, এলকে (Elasticsearch/Logstash/Kibana) বা লোকিতে সরাসরি ক্যোয়ারী করার জন্য স্ট্রাকচার্ড JSON লগ তৈরি করা।

2. **Prometheus Metrics**:
   - `http_requests_total` কাউন্টার ও হিস্টোগ্রাম দিয়ে অ্যাপ্লিকেশনের পারফরম্যান্স পরিমাপ।

---

## 🚀 How to Run (কোড চালনার নিয়ম)

```bash
go run ./projects/08_observability/main.go
```
