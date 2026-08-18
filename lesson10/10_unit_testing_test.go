package lesson10

import (
	"errors"
	"testing"
)

// ১. টেস্টের জন্য মূল ফাংশন (Function to be tested)
func Add(a, b int) int {
	return a + b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// ২. সাধারণ ইউনিট টেস্ট (Simple Unit Test)
func TestAdd(t *testing.T) {
	got := Add(10, 20)
	want := 30

	if got != want {
		t.Errorf("Add(10, 20) = %d; want %d", got, want)
	}
}

// ৩. টেবিল ড্রাইভেন টেস্ট (Table-Driven Test) - Go-এর বেস্ট প্র্যাকটিস
func TestDivideTableDriven(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		wantResult  float64
		expectError bool
	}{
		{name: "Valid division", a: 10, b: 2, wantResult: 5, expectError: false},
		{name: "Division by zero", a: 10, b: 0, wantResult: 0, expectError: true},
		{name: "Negative division", a: -12, b: 3, wantResult: -4, expectError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%f, %f) expected error, but got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%f, %f) unexpected error: %v", tt.a, tt.b, err)
				}
				if got != tt.wantResult {
					t.Errorf("Divide(%f, %f) = %f; want %f", tt.a, tt.b, got, tt.wantResult)
				}
			}
		})
	}
}

// ৪. বেঞ্চমার্ক টেস্ট (Benchmark Test - পারফরম্যান্স পরীক্ষা)
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}
