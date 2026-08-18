package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// বইয়ের ডেটা মডেল (Book Struct)
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// ইন-মেমোরি ডাটাবেস এবং মিউটেক্স (In-Memory DB with Mutex safety)
var (
	books = []Book{
		{ID: 1, Title: "গো প্রোগ্রামিং শেখার সহজ পাঠ", Author: "গৌতম রায়", Price: 350},
		{ID: 2, Title: "কনকারেন্সি এবং গোরুটিন মাস্টারক্লাস", Author: "সুজন রায়", Price: 450},
	}
	mu sync.Mutex
)

// সকল বই প্রাপ্তি এবং নতুন বই যুক্ত করার হ্যান্ডলার
func booksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case http.MethodGet:
		// কোয়েরি প্যারামিটার পরীক্ষা (?id=1)
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, `{"error": "অবৈধ বই আইডি"}`, http.StatusBadRequest)
				return
			}

			mu.Lock()
			defer mu.Unlock()
			for _, b := range books {
				if b.ID == id {
					json.NewEncoder(w).Encode(b)
					return
				}
			}
			http.Error(w, `{"error": "বইটি পাওয়া যায়নি"}`, http.StatusNotFound)
			return
		}

		// সকল বই রিটার্ন করা
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var newBook Book
		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
			http.Error(w, `{"error": "অবৈধ ডেটা ফরমেট"}`, http.StatusBadRequest)
			return
		}

		mu.Lock()
		newBook.ID = len(books) + 1
		books = append(books, newBook)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "নতুন বই সফলভাবে যুক্ত হয়েছে!",
			"book":    newBook,
		})

	default:
		http.Error(w, `{"error": "মেথড সাপোর্ট করে না"}`, http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/books", booksHandler)

	fmt.Println("==================================================")
	fmt.Println(" 🚀 গো (Golang) REST API সার্ভার চালু হয়েছে! ")
	fmt.Println(" 📍 ইউআরএল: http://localhost:8080/api/books")
	fmt.Println("==================================================")

	// সার্ভার লিসেন করা
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("সার্ভার ত্রুটি:", err)
	}
}
