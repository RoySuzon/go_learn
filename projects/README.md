# 🏗️ Real-World Systems Engineering Projects Index

স্বাগতম **গো (Golang) রিয়েল-ওয়ার্ল্ড প্রজেক্টস মডিউল**-এ। এই ফোল্ডারে ৮টি ওয়েব অ্যাপ্লিকেশন, মাইক্রোসার্ভিস, ডিস্ট্রিবিউটেড সিস্টেম এবং সিস্টেম ডিজাইন প্যাটার্নের প্রজেক্ট বাংলায় বিস্তারিত নির্দেশিকাসহ যুক্ত করা হয়েছে।

---

## 📌 Projects Directory Matrix (প্রজেক্ট সূচিপত্র)

| # | Project Name | Description / Tech Stack | Bangla Guide | Source Code | Command to Run |
|---|--------------|--------------------------|--------------|-------------|----------------|
| **01** | **RESTful API** | JWT Bearer Auth + GORM ORM + PostgreSQL | [বাংলা গাইড](01_rest_api/README.md) | [`01_rest_api/`](01_rest_api/main.go) | `go run ./projects/01_rest_api/main.go` |
| **02** | **CLI Todo App** | Command-line Flags & JSON File Storage | [বাংলা গাইড](02_cli_todo/README.md) | [`02_cli_todo/`](02_cli_todo/main.go) | `go run ./projects/02_cli_todo/main.go -list` |
| **03** | **gRPC Microservice** | HTTP/2 Binary ProtoBuf Server & Client | [বাংলা গাইড](03_grpc_service/README.md) | [`03_grpc_service/`](03_grpc_service/) | `go run ./projects/03_grpc_service/server/main.go` |
| **04** | **Redis Caching** | Cache-Aside Pattern with TTL Expiration | [বাংলা গাইড](04_redis_cache/README.md) | [`04_redis_cache/`](04_redis_cache/main.go) | `go run ./projects/04_redis_cache/main.go` |
| **05** | **Event Queue Broker** | Async Message Queue, Worker Pool & DLQ | [বাংলা গাইড](05_event_queue/README.md) | [`05_event_queue/`](05_event_queue/main.go) | `go run ./projects/05_event_queue/main.go` |
| **06** | **Advanced Database** | DB Connection Pooling (`MaxOpenConns`) & ACID Tx | [বাংলা গাইড](06_database_advanced/README.md) | [`06_database_advanced/`](06_database_advanced/main.go) | `go run ./projects/06_database_advanced/main.go` |
| **07** | **System Design** | Token Bucket Rate Limiter & Circuit Breaker | [বাংলা গাইড](07_system_design/README.md) | [`07_system_design/`](07_system_design/main.go) | `go run ./projects/07_system_design/main.go` |
| **08** | **Observability** | Structured JSON (`log/slog`) & Prometheus Metrics | [বাংলা গাইড](08_observability/README.md) | [`08_observability/`](08_observability/main.go) | `go run ./projects/08_observability/main.go` |

---

## 🧪 Test Commands
```bash
go test -v ./...
```
