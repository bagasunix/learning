package factorial

// =============================================================================
// FACTORIAL
// =============================================================================
// Tingkat   : Mudah
// Konsep    : Loop, rekursi, matematika dasar
// Complexity: O(n) time, O(1) space (iterative) / O(n) stack (recursive)
// =============================================================================
//
// SOAL:
//   Hitung n! (n faktorial) = n * (n-1) * (n-2) * ... * 2 * 1
//   (LeetCode tidak punya soal ini karena terlalu mudah, tapi sering muncul di interview)
//
// CONTOH:
//   factorial(0)  -> 1    (by definition, 0! = 1)
//   factorial(1)  -> 1
//   factorial(5)  -> 120  (5*4*3*2*1)
//   factorial(10) -> 3628800
//
// EDGE CASES yang harus disebut:
//   - n = 0  -> 1  (definisi matematika, bukan 0!)
//   - n = 1  -> 1
//   - n < 0  -> tidak terdefinisi, bisa return -1 atau error
//   - n > 20 -> overflow int64 di Go (20! = 2.4 * 10^18 masih muat, 21! tidak)
//
// ADA DUA IMPLEMENTASI:
//
//   1. factorial (iterative):
//      - Loop dari 2 sampai n, kalikan result
//      - O(n) time, O(1) space ← LEBIH BAIK
//
//   2. factorialRecursive:
//      - Base case: n <= 1 return 1
//      - Rekursi: n * factorial(n-1)
//      - O(n) time, O(n) stack space
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Iterative lebih efisien karena tidak ada overhead stack call.
//    Base case n=0 dan n=1 keduanya return 1.
//    Perlu ingatkan bahwa n! overflow di int untuk n > 20."
//
// JEBAKAN UMUM:
//   - Lupa 0! = 1 (bukan 0)
//   - Tidak handle n < 0
//   - Loop mulai dari 1 bukan 2 (tidak salah, tapi tidak efisien)
// =============================================================================

// factorial: iterative — O(n) time, O(1) space
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// factorialRecursive: O(n) time, O(n) stack
func factorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}
