# 🌐 Project 01: RESTful Web API Server (JWT Auth + GORM PostgreSQL)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
একটি রিয়েল-ওয়ার্ল্ড **RESTful HTTP Web API Server** যা **JWT Bearer Token Authentication**, পাসওয়ার্ড হ্যাশিং এবং **GORM ORM + PostgreSQL** ডাটাবেস পারসিস্টেন্স সমর্থন করে।

---

## 🔑 Key Features (প্রধান সুবিধাসমূহ)

1. **JWT Bearer Token Auth**: ইউজারের পরিচয় প্রমাণের জন্য ডিজিটাল টোকেন তৈরি ও ভ্যালিডেশন।
2. **Password Hashing**: কাঁচা পাসওয়ার্ডের বদলে হ্যাশ মান সংরক্ষণ।
3. **GORM ORM & PostgreSQL**: ডাটাবেস টেবিল অটো-মাইগ্রেশন (`db.AutoMigrate`).
4. **Standalone Fallback**: PostgreSQL কানেক্ট না থাকলেও অ্যাপ ইন-মেমোরি মোডে মসৃণভাবে চলবে।

---

## 📌 API Endpoints Matrix (এপিআই এন্ডপয়েন্টসমূহ)

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/register` | নতুন ইউজার রেজিস্ট্রেশন করা (`{"username": "goutom", "password": "123"}`) | ❌ No |
| `POST` | `/api/login` | লগইন করে JWT Bearer Token প্রাপ্তি | ❌ No |
| `GET` | `/api/books` | বইয়ের তালিকা পড়া | ❌ No |
| `GET` | `/api/profile` | **সুরক্ষিত ইউজার প্রোফাইল** (`Authorization: Bearer <token>`) | ✅ Yes |

---

## 🚀 How to Run & Test (কোড চালনার নিয়ম)

### 1. Run Server Standalone:
```bash
go run ./projects/01_rest_api/main.go
```

### 2. Run with Docker Compose (Go App + PostgreSQL DB):
```bash
docker compose up -d
```

### 3. Run Unit Tests:
```bash
go test -v ./projects/01_rest_api/...
```
