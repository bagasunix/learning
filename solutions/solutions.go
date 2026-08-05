package solutions

// Package solutions berisi jawaban referensi untuk semua soal.
// File ini berisi implementasi lengkap yang sudah teruji.

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// ============== FIZZBUZZ ==============

func Fizzbuzz(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// ============== FIBONACCI ==============

func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	result := make([]int, n)
	result[0], result[1] = 0, 1
	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result
}

func FibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

func FibMemoized(n int) int {
	memo := make(map[int]int)
	var helper func(int) int
	helper = func(k int) int {
		if k < 2 {
			return k
		}
		if v, ok := memo[k]; ok {
			return v
		}
		memo[k] = helper(k-1) + helper(k-2)
		return memo[k]
	}
	return helper(n)
}

// ============== PALINDROME ==============

func cleanString(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func IsPalindrome(s string) bool {
	cleaned := cleanString(s)
	runes := []rune(cleaned)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindromeTwoPointer(s string) bool {
	cleaned := cleanString(s)
	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// ============== ANAGRAM ==============

func cleanAlpha(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func IsAnagram(a, b string) bool {
	aClean, bClean := cleanAlpha(a), cleanAlpha(b)
	if len(aClean) != len(bClean) {
		return false
	}
	aRunes, bRunes := []rune(aClean), []rune(bClean)
	sort.Slice(aRunes, func(i, j int) bool { return aRunes[i] < aRunes[j] })
	sort.Slice(bRunes, func(i, j int) bool { return bRunes[i] < bRunes[j] })
	return string(aRunes) == string(bRunes)
}

func IsAnagramCounter(a, b string) bool {
	aClean, bClean := cleanAlpha(a), cleanAlpha(b)
	if len(aClean) != len(bClean) {
		return false
	}
	counter := make(map[rune]int)
	for _, c := range aClean {
		counter[c]++
	}
	for _, c := range bClean {
		counter[c]--
		if counter[c] < 0 {
			return false
		}
	}
	return true
}

// ============== TWO SUM ==============

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumHashmap(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

// ============== REVERSE STRING ==============

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseStringRecursive(s string) string {
	if len(s) <= 1 {
		return s
	}
	runes := []rune(s)
	return ReverseStringRecursive(string(runes[1:])) + string(runes[0])
}

// ============== VALID PARENTHESES ==============

func IsValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// ============== FACTORIAL ==============

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func FactorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

// ============== PRIME NUMBER ==============

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func SievePrimes(n int) []int {
	if n < 2 {
		return []int{}
	}
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}
	var result []int
	for i := 2; i <= n; i++ {
		if sieve[i] {
			result = append(result, i)
		}
	}
	return result
}

// ============== COUNT VOWELS ==============

func CountVowels(s string) int {
	vowels := "aiueo"
	count := 0
	for _, c := range strings.ToLower(s) {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}
	return count
}

func CountVowelsDetail(s string) map[string]int {
	result := map[string]int{"a": 0, "i": 0, "u": 0, "e": 0, "o": 0}
	for _, c := range strings.ToLower(s) {
		key := string(c)
		if _, ok := result[key]; ok {
			result[key]++
		}
	}
	return result
}

// ============== BINARY SEARCH ==============

func BinarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursive(nums []int, target int) int {
	var helper func(lo, hi int) int
	helper = func(lo, hi int) int {
		if lo > hi {
			return -1
		}
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			return helper(mid+1, hi)
		}
		return helper(lo, mid-1)
	}
	return helper(0, len(nums)-1)
}

// ============== MERGE SORTED ==============

func MergeSorted(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// ============== MAX SUBARRAY (Kadane) ==============

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, curSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

// ============== FIRST UNIQUE CHAR ==============

func FirstUniqChar(s string) int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}
	for i, c := range s {
		if counts[c] == 1 {
			return i
		}
	}
	return -1
}

// ============== LINKED LIST CYCLE ==============

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// ============== STACK QUEUE ==============

type MyQueue struct {
	in  []int
	out []int
}

func NewQueue() *MyQueue {
	return &MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
	q.shift()
	val := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return val
}

func (q *MyQueue) Peek() int {
	q.shift()
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func (q *MyQueue) shift() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}
}

// ============== CLIMB STAIRS ==============

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// ============== MAJORITY ELEMENT ==============

func MajorityElement(nums []int) int {
	candidate, count := nums[0], 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// ============== ROTATE ARRAY ==============

func Rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}
	reverseSlice(nums, 0, n-1)
	reverseSlice(nums, 0, k-1)
	reverseSlice(nums, k, n-1)
}

func reverseSlice(nums []int, lo, hi int) {
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
}

// ============== STRING COMPRESSION ==============

func Compress(s string) string {
	if len(s) == 0 {
		return ""
	}
	var b strings.Builder
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			b.WriteByte(s[i-1])
			b.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	b.WriteByte(s[len(s)-1])
	b.WriteString(strconv.Itoa(count))
	return b.String()
}
