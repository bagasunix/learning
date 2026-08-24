package interviewprep

import (
	"testing"
)

// =============================================================================
// TEST — INTERVIEW PREP
// Jalankan: go test ./interview_prep/... -v
// =============================================================================

// --- FizzBuzz ---
func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	want := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	if len(got) != len(want) {
		t.Fatalf("FizzBuzz(15) len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("FizzBuzz index %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFizzBuzzEdge(t *testing.T) {
	if len(FizzBuzz(0)) != 0 {
		t.Error("FizzBuzz(0) harus return slice kosong")
	}
	if FizzBuzz(1)[0] != "1" {
		t.Error("FizzBuzz(1)[0] harus '1'")
	}
}

// --- Fibonacci ---
func TestFibonacci(t *testing.T) {
	tests := []struct{ n int; want []int }{
		{0, []int{}},
		{1, []int{0}},
		{5, []int{0, 1, 1, 2, 3}},
		{8, []int{0, 1, 1, 2, 3, 5, 8, 13}},
	}
	for _, tt := range tests {
		got := Fibonacci(tt.n)
		if len(got) != len(tt.want) {
			t.Errorf("Fibonacci(%d) len=%d, want %d", tt.n, len(got), len(tt.want))
			continue
		}
		for i := range tt.want {
			if got[i] != tt.want[i] {
				t.Errorf("Fibonacci(%d)[%d] = %d, want %d", tt.n, i, got[i], tt.want[i])
			}
		}
	}
}

func TestFib(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {6, 8}, {10, 55},
	}
	for _, tt := range tests {
		if got := Fib(tt.n); got != tt.want {
			t.Errorf("Fib(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

// --- Palindrome ---
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"A man a plan a canal Panama", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"race a car", false},
		{"Was it a car or a cat I saw?", true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.input); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

// --- Anagram ---
func TestIsAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"listen", "silent", true},
		{"hello", "world", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "a", false},
		{"Astronomer", "moonStarer", true}, // a-s-t-r-o-n-o-m-e-r == m-o-o-n-s-t-a-r-e-r
		{"rat", "car", false},
		{"anagram", "nagaram", true},
	}
	for _, tt := range tests {
		if got := IsAnagram(tt.a, tt.b); got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

// --- Reverse Words ---
func TestReverseWords(t *testing.T) {
	tests := []struct{ input, want string }{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"hello", "hello"},
		{"   ", ""},
		{"", ""},
	}
	for _, tt := range tests {
		if got := ReverseWords(tt.input); got != tt.want {
			t.Errorf("ReverseWords(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

// --- Two Sum ---
func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)
		if got[0] != tt.want[0] || got[1] != tt.want[1] {
			t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

// --- Max Subarray ---
func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1, -2, -3}, -1},
		{[]int{-2, -1}, -1},
	}
	for _, tt := range tests {
		if got := MaxSubArray(tt.nums); got != tt.want {
			t.Errorf("MaxSubArray(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

// --- Rotate Array ---
func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3}, 4, []int{3, 1, 2}}, // k > n
		{[]int{1}, 5, []int{1}},              // single element
	}
	for _, tt := range tests {
		got := make([]int, len(tt.nums))
		copy(got, tt.nums)
		Rotate(got, tt.k)
		for i := range tt.want {
			if got[i] != tt.want[i] {
				t.Errorf("Rotate(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
				break
			}
		}
	}
}

// --- Valid Parentheses ---
func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"([])", true},
		{"{[]}", true},
		{"(]", false},
		{"([)]", false},
		{"", true},
		{"(", false},
		{")", false},
		{"(((",false},
	}
	for _, tt := range tests {
		if got := IsValid(tt.input); got != tt.want {
			t.Errorf("IsValid(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

// --- Binary Search ---
func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 7, 9, 11}, 7, 3},
		{[]int{1, 3, 5, 7, 9, 11}, 4, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 2}, 2, 1},
		{[]int{}, 5, -1},
	}
	for _, tt := range tests {
		if got := BinarySearch(tt.nums, tt.target); got != tt.want {
			t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}

// --- Prime Number ---
func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false}, {1, false}, {2, true}, {3, true},
		{4, false}, {7, true}, {9, false}, {17, true}, {49, false},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

// --- Reverse Integer ---
func TestReverseInt(t *testing.T) {
	tests := []struct{ input, want int }{
		{123, 321}, {-123, -321}, {120, 21},
		{0, 0}, {1534236469, 0}, {-1534236469, 0},
	}
	for _, tt := range tests {
		if got := ReverseInt(tt.input); got != tt.want {
			t.Errorf("ReverseInt(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

// --- Factorial ---
func TestFactorial(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {5, 120}, {10, 3628800},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
