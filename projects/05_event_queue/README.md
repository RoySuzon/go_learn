# 📬 Event-Driven Architecture & Message Queue Broker

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
হাই-পারফরম্যান্স ব্যাকএন্ড সিস্টেমে ইউজার রিকোয়েস্ট আটকে না রেখে **অ্যাসিনক্রোনাস কাজ** (ইমেইল পাঠানো, ইমেজ প্রসেসিং, পেমেন্ট প্রসেসিং) সম্পন্ন করতে **Event-Driven Architecture (Kafka / RabbitMQ)** ব্যবহার করা হয়।

---

## 🎨 Message Queue Flow Diagram

```text
┌──────────────┐    Event     ┌────────────────┐   Worker 1   ┌───────────────┐
│   Producer   ├─────────────►│ Message Queue  ├─────────────►│  Consumer 1   │
└──────────────┘              └───────┬────────┘              └───────────────┘
                                      │            Worker 2   ┌───────────────┐
                                      ├──────────────────────►│  Consumer 2   │
                                      │                       └───────────────┘
                                      ▼ (On Failure)
                              ┌───────────────┐
                              │  Dead Letter  │
                              │  Queue (DLQ)  │
                              └───────────────┘
```

---

## 🚀 How to Run (কোড চালনার নিয়ম)

```bash
go run ./projects/05_event_queue/main.go
```
