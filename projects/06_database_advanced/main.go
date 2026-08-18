package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Simulated SQL Database Transaction Demo
func executeBankTransfer(tx *sql.Tx, fromAcc, toAcc int, amount float64) error {
	// 1. Deduct from Sender
	fmt.Printf("💳 Deducting $%.2f from Account #%d...\n", amount, fromAcc)

	// 2. Add to Receiver
	fmt.Printf("💰 Adding $%.2f to Account #%d...\n", amount, toAcc)

	// Simulate success
	return nil
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🚀 Advanced SQL Database Engineering & Transactions")
	fmt.Println("==================================================")

	fmt.Println("\n১. DB Connection Pool Configuration (প্রোডাকশন বেস্ট প্র্যাকটিস):")
	fmt.Println("   - SetMaxOpenConns(25)     : সর্বোচ্চ খোলা রাখতে দেওয়া কানেকশন")
	fmt.Println("   - SetMaxIdleConns(5)      : আইডল অবস্থায় প্রস্তুত থাকা কানেকশন")
	fmt.Println("   - SetConnMaxLifetime(5m)  : কানেকশনের সর্বোচ্চ মেয়াদকাল")

	fmt.Println("\n২. ACID Database Transaction Simulation (Bank Transfer):")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Transaction Rollback / Commit Pattern Demo
	err := func() error {
		// Mock Tx
		fmt.Println("   --> Beginning Transaction (BEGIN)...")
		if err := executeBankTransfer(nil, 101, 102, 250.00); err != nil {
			fmt.Println("   ❌ Transaction Failed! Rolling Back (ROLLBACK)...")
			return err
		}
		fmt.Println("   ✅ Transaction Succeeded! Committing Changes (COMMIT)...")
		return nil
	}()

	if err == nil {
		fmt.Println("\n🎉 Bank Transfer Transaction Completed Successfully with ACID guarantees!")
	}
	_ = ctx
}
