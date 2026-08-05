package fibonacci

func fibonacci(n int) []int {
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

func fibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibMemoized(n int) int {
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
