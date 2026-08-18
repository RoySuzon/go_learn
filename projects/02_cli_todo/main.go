package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

const filename = "todos.json"

// ফাইল থেকে টুডু তালিকা পড়া
func loadTodos() ([]Todo, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Todo{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	return todos, err
}

// ফাইলে টুডু তালিকা সংরক্ষণ
func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func main() {
	// কমান্ড লাইন ফ্ল্যাগসমূহ (Command-line flags)
	addFlag := flag.String("add", "", "নতুন টাস্ক যোগ করুন (যেমন: -add \"Go শেখা\")")
	listFlag := flag.Bool("list", false, "সমস্ত টাস্ক তালিকা দেখুন")
	completeFlag := flag.Int("complete", 0, "টাস্ক সম্পন্ন হিসেবে চিহ্নিত করুন (যেমন: -complete 1)")

	flag.Parse()

	todos, err := loadTodos()
	if err != nil {
		fmt.Println("ত্রুটি: টুডু তালিকা লোড করা যায়নি -", err)
		return
	}

	// ১. নতুন টাস্ক যোগ
	if *addFlag != "" {
		newTodo := Todo{
			ID:        len(todos) + 1,
			Task:      *addFlag,
			Completed: false,
		}
		todos = append(todos, newTodo)
		if err := saveTodos(todos); err != nil {
			fmt.Println("ত্রুটি: টাস্ক সংরক্ষণ করা যায়নি -", err)
			return
		}
		fmt.Printf("✅ টাস্ক সফলভাবে যুক্ত হয়েছে: [%d] %s\n", newTodo.ID, newTodo.Task)
		return
	}

	// ২. সম্পন্ন মার্ক করা
	if *completeFlag > 0 {
		found := false
		for i, t := range todos {
			if t.ID == *completeFlag {
				todos[i].Completed = true
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("❌ %d আইডি চিহ্নিত কোনো টাস্ক পাওয়া যায়নি!\n", *completeFlag)
			return
		}
		saveTodos(todos)
		fmt.Printf("🎉 টাস্ক #%d সম্পন্ন হিসেবে চিহ্নিত করা হয়েছে!\n", *completeFlag)
		return
	}

	// ৩. টাস্ক তালিকা প্রদর্শন
	if *listFlag || len(os.Args) == 1 {
		fmt.Println("==========================================")
		fmt.Println("       📝 আপনার টুডু তালিকা (CLI Todo)     ")
		fmt.Println("==========================================")

		if len(todos) == 0 {
			fmt.Println("কোনো টাস্ক পাওয়া যায়নি! নতুন টাস্ক যোগ করতে টাইপ করুন:")
			fmt.Println("  go run main.go -add \"আপনার কাজ\"")
			return
		}

		for _, t := range todos {
			status := "❌ অসম্পূর্ণ"
			if t.Completed {
				status = "✅ সম্পন্ন"
			}
			fmt.Printf("[%d] %-30s | %s\n", t.ID, t.Task, status)
		}
		fmt.Println("==========================================")
	}
}
