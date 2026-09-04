package fibonacci

// =============================================================================
// FIBONACCI
// =============================================================================
// Tingkat   : Mudah
// Konsep    : Rekursi, iterasi, memoization, dynamic programming
// =============================================================================
//
// SOAL:
//   Ada 3 variasi yang sering muncul di interview:
//
//   VARIASI 1 — fibonacci(n): return slice n angka pertama deret fibonacci
//     fibonacci(5) -> [0, 1, 1, 2, 3]
//     fibonacci(1) -> [0]
//     fibonacci(0) -> []
//
//   VARIASI 2 — fib(n): return nilai fibonacci ke-n (0-indexed)
//     fib(0) -> 0
//     fib(1) -> 1
//     fib(6) -> 8   (0,1,1,2,3,5,8)
//
//   VARIASI 3 — fibRecursive(n): sama seperti fib(n) tapi rekursif
//     (interviewer sering minta ini untuk test pemahaman rekursi)
//
// DERET: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
//        Setiap angka = dua angka sebelumnya
//
// EDGE CASES yang harus disebut:
//   - n = 0  -> 0 (atau slice kosong tergantung variasi)
//   - n = 1  -> 1
//   - n < 0  -> tanyakan ke interviewer
//
// PERBANDINGAN APPROACH (wajib bisa explain):
//
//   1. Iterative (fib)        : O(n) time, O(1) space  ← TERBAIK
//   2. Memoized (fibMemoized) : O(n) time, O(n) space  ← bagus untuk rekursi
//   3. Recursive (fibRecursive): O(2^n) time, O(n) space ← LAMBAT, jangan pakai production
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Rekursi biasa O(2^n) karena hitung ulang subproblem yang sama.
//    Dengan memoization jadi O(n) karena tiap nilai cuma dihitung sekali.
//    Iterative paling efisien: O(n) time, O(1) space, cukup simpan dua variabel."
//
// JEBAKAN UMUM:
//   - Lupa base case n=0 dan n=1 di rekursi -> stack overflow
//   - fibRecursive(50) akan sangat lambat tanpa memoization
// =============================================================================

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

// fibRecursive: O(2^n) — untuk demo rekursi saja, bukan production
func fibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// fibMemoized: O(n) time, O(n) space — rekursi dengan cache
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

// fib: O(n) time, O(1) space — PALING OPTIMAL
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
