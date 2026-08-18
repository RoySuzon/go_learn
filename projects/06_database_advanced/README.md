# 🗄️ Advanced Database Engineering & SQL Optimization

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
প্রোডাকশন সিস্টেমে ডাটাবেসের কর্মক্ষমতা বৃদ্ধি ও ডাটা সুরক্ষা নিশ্চিত করতে **Connection Pooling**, **ACID Transactions**, এবং **$B\text{-Tree}$ Indexing** অপরিহার্য।

---

## ⚙️ Production Connection Pool Settings

```go
db.SetMaxOpenConns(25)                  // Maximum active connections
db.SetMaxIdleConns(5)                   // Maximum idle connections in pool
db.SetConnMaxLifetime(5 * time.Minute)  // Maximum lifetime of a connection
```

---

## 🔒 ACID Transaction Pattern

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil {
    return err
}
defer tx.Rollback() // Rolled back if not committed

// Execute SQL statements inside transaction
if err := updateSender(tx); err != nil { return err }
if err := updateReceiver(tx); err != nil { return err }

return tx.Commit() // Commit transaction
```

---

## 🚀 How to Run (কোড চালনার নিয়ম)

```bash
go run ./projects/06_database_advanced/main.go
```
