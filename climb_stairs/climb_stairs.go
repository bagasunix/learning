package climbstairs

// Climbing Stairs
// ---------------
// Ada n tangga. Tiap langkah bisa naik 1 atau 2 anak tangga.
// Berapa cara untuk mencapai tangga ke-n? (LeetCode #70)
//
// Contoh:
//
//	climbStairs(2) -> 2  (1+1, 2)
//	climbStairs(3) -> 3  (1+1+1, 1+2, 2+1)
//	climbStairs(5) -> 8
//
// Hint: Ini pada dasarnya fibonacci. Gunakan DP atau memoization.
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
