package solutions

import (
	"reflect"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	if !reflect.DeepEqual(Fizzbuzz(15), expected) {
		t.Errorf("got %v, want %v", Fizzbuzz(15), expected)
	}
}

func TestFibonacci(t *testing.T) {
	if !reflect.DeepEqual(Fibonacci(7), []int{0, 1, 1, 2, 3, 5, 8}) {
		t.Errorf("got %v", Fibonacci(7))
	}
}

func TestFibRecursive(t *testing.T) {
	if FibRecursive(10) != 55 {
		t.Errorf("got %d", FibRecursive(10))
	}
}

func TestFibMemoized(t *testing.T) {
	if FibMemoized(50) != 12586269025 {
		t.Errorf("got %d", FibMemoized(50))
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("racecar") {
		t.Error("racecar should be palindrome")
	}
	if IsPalindrome("hello") {
		t.Error("hello should not be palindrome")
	}
	if !IsPalindrome("A man a plan a canal Panama") {
		t.Error("should be palindrome")
	}
}

func TestIsAnagram(t *testing.T) {
	if !IsAnagram("listen", "silent") {
		t.Error("should be anagram")
	}
	if IsAnagram("hello", "world") {
		t.Error("should not be anagram")
	}
	if !IsAnagramCounter("Dormitory", "dirty room") {
		t.Error("should be anagram")
	}
}

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(TwoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}) {
		t.Error("failed")
	}
	if !reflect.DeepEqual(TwoSumHashmap([]int{3, 2, 4}, 6), []int{1, 2}) {
		t.Error("failed")
	}
}

func TestReverseString(t *testing.T) {
	if ReverseString("hello") != "olleh" {
		t.Error("failed")
	}
	if ReverseStringRecursive("hello") != "olleh" {
		t.Error("failed")
	}
}

func TestIsValid(t *testing.T) {
	if !IsValid("()[]{}") {
		t.Error("should be valid")
	}
	if IsValid("([)]") {
		t.Error("should be invalid")
	}
	if !IsValid("{[]}") {
		t.Error("should be valid")
	}
}

func TestFactorial(t *testing.T) {
	if Factorial(5) != 120 {
		t.Error("failed")
	}
	if FactorialRecursive(5) != 120 {
		t.Error("failed")
	}
}

func TestIsPrime(t *testing.T) {
	if !IsPrime(7) {
		t.Error("7 should be prime")
	}
	if IsPrime(10) {
		t.Error("10 should not be prime")
	}
	if !reflect.DeepEqual(SievePrimes(20), []int{2, 3, 5, 7, 11, 13, 17, 19}) {
		t.Errorf("got %v", SievePrimes(20))
	}
}

func TestCountVowels(t *testing.T) {
	if CountVowels("hello") != 2 {
		t.Error("failed")
	}
	detail := CountVowelsDetail("hello")
	if detail["e"] != 1 || detail["o"] != 1 {
		t.Errorf("got %v", detail)
	}
}

// ============== BINARY SEARCH ==============

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if BinarySearch(nums, 7) != 3 {
		t.Error("failed")
	}
	if BinarySearch(nums, 4) != -1 {
		t.Error("failed")
	}
	if BinarySearch(nums, 1) != 0 {
		t.Error("failed")
	}
	if BinarySearch(nums, 11) != 5 {
		t.Error("failed")
	}
	if BinarySearch([]int{}, 5) != -1 {
		t.Error("failed")
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if BinarySearchRecursive(nums, 7) != 3 {
		t.Error("failed")
	}
	if BinarySearchRecursive(nums, 4) != -1 {
		t.Error("failed")
	}
	if BinarySearchRecursive([]int{}, 5) != -1 {
		t.Error("failed")
	}
}

// ============== MERGE SORTED ==============

func TestMergeSorted(t *testing.T) {
	if !reflect.DeepEqual(MergeSorted([]int{1, 3, 5}, []int{2, 4, 6}), []int{1, 2, 3, 4, 5, 6}) {
		t.Error("failed")
	}
	if !reflect.DeepEqual(MergeSorted([]int{1, 2, 3}, []int{}), []int{1, 2, 3}) {
		t.Error("failed")
	}
	if !reflect.DeepEqual(MergeSorted([]int{}, []int{}), []int{}) {
		t.Error("failed")
	}
}

// ============== MAX SUBARRAY ==============

func TestMaxSubArray(t *testing.T) {
	if MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("failed")
	}
	if MaxSubArray([]int{1}) != 1 {
		t.Error("failed")
	}
	if MaxSubArray([]int{-3, -1, -2}) != -1 {
		t.Error("failed")
	}
}

// ============== FIRST UNIQUE CHAR ==============

func TestFirstUniqChar(t *testing.T) {
	if FirstUniqChar("leetcode") != 0 {
		t.Error("failed")
	}
	if FirstUniqChar("loveleetcode") != 2 {
		t.Error("failed")
	}
	if FirstUniqChar("aabb") != -1 {
		t.Error("failed")
	}
}

// ============== LINKED LIST CYCLE ==============

func TestHasCycle(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	n3.Next = n2
	if !HasCycle(n1) {
		t.Error("should have cycle")
	}

	n3.Next = nil
	if HasCycle(n1) {
		t.Error("should not have cycle")
	}

	if HasCycle(nil) {
		t.Error("nil should not have cycle")
	}
}

// ============== STACK QUEUE ==============

func TestMyQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 {
		t.Error("peek failed")
	}
	if q.Pop() != 1 {
		t.Error("pop failed")
	}
	if q.Empty() {
		t.Error("should not be empty")
	}
	if q.Pop() != 2 {
		t.Error("pop failed")
	}
	if !q.Empty() {
		t.Error("should be empty")
	}
}

// ============== CLIMB STAIRS ==============

func TestClimbStairs(t *testing.T) {
	if ClimbStairs(1) != 1 {
		t.Error("failed")
	}
	if ClimbStairs(2) != 2 {
		t.Error("failed")
	}
	if ClimbStairs(3) != 3 {
		t.Error("failed")
	}
	if ClimbStairs(5) != 8 {
		t.Error("failed")
	}
	if ClimbStairs(45) != 1836311903 {
		t.Error("failed")
	}
}

// ============== MAJORITY ELEMENT ==============

func TestMajorityElement(t *testing.T) {
	if MajorityElement([]int{3, 2, 3}) != 3 {
		t.Error("failed")
	}
	if MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}) != 2 {
		t.Error("failed")
	}
	if MajorityElement([]int{1}) != 1 {
		t.Error("failed")
	}
}

// ============== ROTATE ARRAY ==============

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(nums, 3)
	if !reflect.DeepEqual(nums, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Errorf("got %v", nums)
	}

	nums2 := []int{-1, -100, 3, 99}
	Rotate(nums2, 2)
	if !reflect.DeepEqual(nums2, []int{3, 99, -1, -100}) {
		t.Errorf("got %v", nums2)
	}

	nums3 := []int{1, 2, 3}
	Rotate(nums3, 3)
	if !reflect.DeepEqual(nums3, []int{1, 2, 3}) {
		t.Errorf("got %v", nums3)
	}
}

// ============== STRING COMPRESSION ==============

func TestCompress(t *testing.T) {
	if Compress("aabcccccaaa") != "a2b1c5a3" {
		t.Error("failed")
	}
	if Compress("abc") != "a1b1c1" {
		t.Error("failed")
	}
	if Compress("") != "" {
		t.Error("failed")
	}
	if Compress("aaaa") != "a4" {
		t.Error("failed")
	}
	if Compress("a") != "a1" {
		t.Error("failed")
	}
}
