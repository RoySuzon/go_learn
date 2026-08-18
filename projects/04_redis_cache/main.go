package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Global Redis Client & Fallback Memory Cache
var (
	ctx        = context.Background()
	rdb        *redis.Client
	localCache = make(map[string]string)
	cacheMu    sync.RWMutex
)

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping Redis server to verify connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("⚠️ Redis Warning:", err)
		fmt.Println("ℹ️ Running with Local In-Memory Fallback Cache.")
		rdb = nil
	} else {
		fmt.Println("✅ Connected to Redis Cache Server successfully!")
	}
}

// GetProduct retrieves product with Cache-Aside Pattern (Cache -> DB -> Cache update)
func GetProduct(id int) (*Product, string) {
	cacheKey := fmt.Sprintf("product:%d", id)

	// 1. Try fetching from Redis or Local Cache
	if rdb != nil {
		val, err := rdb.Get(ctx, cacheKey).Result()
		if err == nil {
			var p Product
			json.Unmarshal([]byte(val), &p)
			return &p, "HIT (Redis)"
		}
	} else {
		cacheMu.RLock()
		val, exists := localCache[cacheKey]
		cacheMu.RUnlock()
		if exists {
			var p Product
			json.Unmarshal([]byte(val), &p)
			return &p, "HIT (Local Cache)"
		}
	}

	// 2. Cache MISS: Simulate DB Fetch (100ms latency)
	time.Sleep(100 * time.Millisecond)
	dbProduct := Product{ID: id, Name: fmt.Sprintf("Laptop Pro #%d", id), Price: 1299.99}

	// 3. Save fetched item to Cache with 10-second TTL
	bytes, _ := json.Marshal(dbProduct)
	if rdb != nil {
		rdb.Set(ctx, cacheKey, string(bytes), 10*time.Second)
	} else {
		cacheMu.Lock()
		localCache[cacheKey] = string(bytes)
		cacheMu.Unlock()
	}

	return &dbProduct, "MISS (Fetched from DB & Cached)"
}

func main() {
	initRedis()

	fmt.Println("==================================================")
	fmt.Println(" 🚀 High-Speed Redis Caching Demo (Cache-Aside)")
	fmt.Println("==================================================")

	// Call 1: Expect Cache MISS (Slow ~100ms)
	start1 := time.Now()
	p1, source1 := GetProduct(101)
	fmt.Printf("1st Call -> Status: %-25s | Time: %v | Data: %+v\n", source1, time.Since(start1), p1)

	// Call 2: Expect Cache HIT (Sub-millisecond ~0ms!)
	start2 := time.Now()
	p2, source2 := GetProduct(101)
	fmt.Printf("2nd Call -> Status: %-25s | Time: %v | Data: %+v\n", source2, time.Since(start2), p2)
}
