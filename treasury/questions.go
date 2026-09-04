package main

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func Question1A(arr []int) int {
	var hasPos, hasNeg bool
	exists := make(map[int]bool, len(arr))
	for _, v := range arr {
		if v > 0 {
			hasPos = true
		} else if v < 0 {
			hasNeg = true
		}
		exists[v] = true
	}
	if hasPos == hasNeg {
		return 0
	}
	if hasPos {
		for i := 1; ; i++ {
			if !exists[i] {
				return i
			}
		}
	}
	for i := -1; ; i-- {
		if !exists[i] {
			return i
		}
	}
}

func Question1B(arr []int) int {
	var hasPos, hasNeg bool
	for _, v := range arr {
		if v > 0 {
			hasPos = true
		} else if v < 0 {
			hasNeg = true
		}
	}
	if hasPos == hasNeg {
		return 0
	}

	result := arr[0]
	for _, v := range arr {
		if hasPos && v < result {
			result = v
		}
		if hasNeg && v > result {
			result = v
		}
	}
	return result
}

func Question2(exclude []int, nums ...int) (int, error) {
	var sb strings.Builder
	for _, n := range nums {
		sb.WriteString(strconv.Itoa(n))
	}
	s := sb.String()

	excSet := make(map[byte]bool)
	for _, e := range exclude {
		if e >= 0 && e <= 9 {
			excSet[byte('0'+e)] = true
		}
	}

	var result []byte
	for i := len(s) - 1; i >= 0; i-- {
		if !excSet[s[i]] {
			result = append(result, s[i])
		}
	}
	if len(result) == 0 {
		return 0, errors.New("not found")
	}
	return strconv.Atoi(string(result))
}

func Question3(numsA, numsB []int) []int {
	bset := make(map[int]bool)
	for _, v := range numsB {
		bset[v] = true
	}
	seen := make(map[int]bool)
	var result []int
	for _, v := range numsA {
		if bset[v] && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	slices.Sort(result)
	slices.Reverse(result)
	return result
}

func Question4(num int) map[int]int {
	if num < 0 {
		num = -num
	}
	result := make(map[int]int)
	for _, r := range strconv.Itoa(num) {
		result[int(r-'0')]++
	}
	return result
}

func Question5(nums ...int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	products := make([]int, (n+1)/2)
	var wg sync.WaitGroup

	for i := range products {
		a, b := nums[i*2], nums[i*2]
		if i*2+1 < n {
			b = nums[i*2+1]
		}
		wg.Go(func() {
			products[i] = a * b
		})
	}
	wg.Wait()

	sum := 0
	for _, p := range products {
		sum += p
	}
	return sum
}
