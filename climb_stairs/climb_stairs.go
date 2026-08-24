package climbstairs

// =============================================================================
// CLIMBING STAIRS
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Dynamic programming, fibonacci variant
// Complexity: O(n) time, O(1) space
// =============================================================================
//
// SOAL:
//   Ada n anak tangga. Setiap langkah bisa naik 1 atau 2 anak tangga.
//   Berapa banyak cara berbeda untuk mencapai tangga ke-n? (LeetCode #70)
//
// CONTOH:
//   climbStairs(1) -> 1   (hanya: [1])
//   climbStairs(2) -> 2   ([1+1], [2])
//   climbStairs(3) -> 3   ([1+1+1], [1+2], [2+1])
//   climbStairs(4) -> 5   ([1+1+1+1], [1+1+2], [1+2+1], [2+1+1], [2+2])
//   climbStairs(5) -> 8
//
// KENAPA INI FIBONACCI?
//   Untuk mencapai tangga ke-n, kamu bisa datang dari:
//   - Tangga ke-(n-1): ambil 1 langkah
//   - Tangga ke-(n-2): ambil 2 langkah
//   Jadi: ways(n) = ways(n-1) + ways(n-2)
//   Base case: ways(1) = 1, ways(2) = 2
//   Ini persis pola fibonacci!
//
//   n:    1  2  3  4  5  6  7
//   ways: 1  2  3  5  8  13 21
//   fib:  1  1  2  3  5  8  13  (fibonacci biasa, tapi offset 1)
//
// EDGE CASES yang harus disebut:
//   - n = 0  -> 1 (ada 1 cara: tidak melakukan apa-apa)
//   - n = 1  -> 1
//   - n = 2  -> 2
//
// CARA KERJA (a, b = prev, curr):
//   n=5: a=1, b=1
//   i=2: a=1, b=1+1=2
//   i=3: a=2, b=1+2=3
//   i=4: a=3, b=2+3=5
//   i=5: a=5, b=3+5=8
//   return b=8 ✓
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Ini fibonacci variant. ways(n) = ways(n-1) + ways(n-2).
//    Tidak perlu array DP, cukup dua variabel: a (prev) dan b (curr).
//    O(n) time, O(1) space."
//
// JEBAKAN UMUM:
//   - Langsung rekursi tanpa memoization -> O(2^n) time limit exceeded
//   - DP array O(n) space -> valid tapi tidak optimal
//   - Base case salah: n<=1 harus return 1 (bukan 0)
// =============================================================================

// climbStairs: DP O(1) space (fibonacci variant)
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
