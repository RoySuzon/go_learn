package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ১. স্ট্রাক্ট এবং কম্পোজিশন (Composition)
type Address struct {
	City    string
	Country string
}

type Student struct {
	ID      int
	Name    string
	Marks   []float64
	Address // Embedded Struct (Composition)
}

// ২. ইন্টারফেস (Interface)
type Evaluator interface {
	CalculateAverage() float64
}

// ৩. মেথড উইথ পয়েন্টার রিসিভার (Pointer Receiver Method)
func (s *Student) AddMark(mark float64) {
	s.Marks = append(s.Marks, mark)
}

// ৪. মেথড উইথ ভ্যালু রিসিভার (Value Receiver Method)
func (s Student) CalculateAverage() float64 {
	if len(s.Marks) == 0 {
		return 0
	}
	total := 0.0
	for _, mark := range s.Marks {
		total += mark
	}
	return total / float64(len(s.Marks))
}

// ৫. এরর হ্যান্ডলিং এবং ফাংশন (Error Handling & Functions)
func DivideMarks(total float64, subjects int) (float64, error) {
	if subjects == 0 {
		return 0, errors.New("বিষয় সংখ্যা শূন্য (0) হতে পারে না")
	}
	return total / float64(subjects), nil
}

// ৬. কনকারেন্সি উইথ থ্রেড সেফটি (Mutex & WaitGroup)
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	fmt.Println("===========================================")
	fmt.Println("   গো (Golang) সম্পূর্ণ বাংলা মাস্টার এক্সাম্পল   ")
	fmt.Println("===========================================")

	// --- ১. ভ্যারিয়েবল এবং স্ট্রাক্ট ইনস্ট্যান্স ---
	student1 := Student{
		ID:   101,
		Name: "গৌতম রায়",
		Address: Address{
			City:    "ঢাকা",
			Country: "বাংলাদেশ",
		},
	}

	// মেথড ব্যবহার করে মার্কস যোগ
	student1.AddMark(85.5)
	student1.AddMark(92.0)
	student1.AddMark(88.0)

	// --- ২. ইন্টারফেস ও মেথড আউটপুট ---
	var eval Evaluator = student1
	avgScore := eval.CalculateAverage()

	fmt.Printf("\n১. ছাত্রের নাম: %s\n", student1.Name)
	fmt.Printf("   ঠিকানা: %s, %s (Promoted Field Access)\n", student1.City, student1.Country)
	fmt.Printf("   গড় নম্বর: %.2f\n", avgScore)

	// --- ৩. এরর হ্যান্ডলিং চেক ---
	fmt.Println("\n২. এরর হ্যান্ডলিং পরীক্ষা:")
	res, err := DivideMarks(265.5, 3)
	if err == nil {
		fmt.Printf("   গড় (ফাংশন): %.2f\n", res)
	}

	_, errZero := DivideMarks(265.5, 0)
	if errZero != nil {
		fmt.Printf("   হ্যান্ডেলকৃত এরর বার্তা: %s\n", errZero)
	}

	// --- ৪. স্ট্যান্ডার্ড লাইব্রেরি প্যাকেজ (strings, strconv, time) ---
	fmt.Println("\n৩. স্ট্যান্ডার্ড লাইব্রেরি ব্যবহার:")
	rawString := "  golang basic course in bangla  "
	cleanStr := strings.TrimSpace(rawString)
	fmt.Println("   ক্লিন স্ট্রিং:", strings.ToUpper(cleanStr))

	strNum := "100"
	parsedInt, _ := strconv.Atoi(strNum)
	fmt.Printf("   পার্সকৃত সংখ্যা + ৫০: %d\n", parsedInt+50)

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("   বর্তমান সময়:", currentTime)
	fmt.Printf("   ১৬-এর বর্গমূল: %.1f\n", math.Sqrt(16))

	// --- ৫. কনকারেন্সি (Goroutines, Channel, WaitGroup, Mutex) ---
	fmt.Println("\n৪. কনকারেন্সি ও থ্রেড-সেফ কাউন্টার:")

	// চ্যানেল উদাহরণ
	msgChan := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		msgChan <- "গোরুটিন থেকে সফলভাবে বার্তা প্রাপ্ত!"
	}()

	msg := <-msgChan
	fmt.Println("   চ্যানেল বার্তা:", msg)

	// WaitGroup ও Mutex উদাহরণ
	var wg sync.WaitGroup
	counter := SafeCounter{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}

	wg.Wait()
	fmt.Println("   কাউন্টার মান (Mutex Safe):", counter.count)

	fmt.Println("\n===========================================")
	fmt.Println("         কোড সফলভাবে সম্পন্ন হয়েছে!        ")
	fmt.Println("===========================================")
}
