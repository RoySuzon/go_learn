# 📝 Project 02: CLI Todo Manager Application

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
একটি কমান্ট-লাইন ইন্টারফেস (CLI) অ্যাপ্লিকেশন যা ফ্ল্যাগ (`flag` package) এবং `JSON` ফাইল পারসিস্টেন্সের মাধ্যমে টুডু তালিকা পরিচালনা করে।

---

## 🔑 Command-Line Flags (কমান্ড-লাইন ফ্ল্যাগসমূহ)

- `-add "টাস্ক নাম"` : নতুন টুডু টাস্ক যোগ করার জন্য।
- `-list` : সমস্ত টুডু টাস্ক দেখার জন্য।
- `-complete <ID>` : নির্দিষ্ট আইডি-র টাস্ক সম্পন্ন হিসেবে চিহ্নিত করার জন্য।

---

## 🚀 How to Run (কোড চালনার নিয়ম)

```bash
# ১. নতুন টাস্ক যোগ করা
go run ./projects/02_cli_todo/main.go -add "Go প্রোগ্রামিং প্র্যাকটিস করা"

# ২. টাস্ক তালিকা দেখা
go run ./projects/02_cli_todo/main.go -list

# ৩. টাস্ক #1 সম্পন্ন মার্ক করা
go run ./projects/02_cli_todo/main.go -complete 1
```
