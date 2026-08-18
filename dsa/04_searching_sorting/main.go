package main

import "fmt"

func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right, equal := []int{}, []int{}, []int{}

	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val > pivot {
			right = append(right, val)
		} else {
			equal = append(equal, val)
		}
	}

	result := append(QuickSort(left), equal...)
	return append(result, QuickSort(right)...)
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 DSA Topic 04: Searching & Sorting Algorithms")
	fmt.Println("==================================================")

	numbers := []int{40, 10, 50, 20, 30}

	fmt.Println("১. অসাজানো ইনপুট অ্যারে:", numbers)
	fmt.Println("   - লিনিয়ার সার্চ (Linear Search - 50):", LinearSearch(numbers, 50))

	fmt.Println("\n২. কুইক সর্ট (Quick Sort) দিয়ে সাজানো হচ্ছে...")
	sorted := QuickSort(numbers)
	fmt.Println("   - সাজানো অ্যারে:", sorted)

	fmt.Println("\n৩. সাজানো অ্যারেতে বাইনারি সার্চ (Binary Search - 30):")
	idx := BinarySearch(sorted, 30)
	fmt.Printf("   - 30 সংখ্যাটি ইনডেক্স %d-এ পাওয়া গেছে!\n", idx)
}
