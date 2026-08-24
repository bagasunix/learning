package primenumber

import "math"

// =============================================================================
// PRIME NUMBER
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Matematika, optimasi loop, Sieve of Eratosthenes
// =============================================================================
//
// ADA DUA SOAL YANG SERING MUNCUL:
//
// SOAL 1 — isPrime(n): cek apakah n adalah bilangan prima
//   isPrime(2)   -> true
//   isPrime(7)   -> true
//   isPrime(1)   -> false  (1 bukan prima by definition)
//   isPrime(9)   -> false  (9 = 3*3)
//   isPrime(0)   -> false
//   isPrime(-5)  -> false
//
// SOAL 2 — sievePrimes(n): cari semua prima dari 2 sampai n
//   sievePrimes(10) -> [2, 3, 5, 7]
//   sievePrimes(20) -> [2, 3, 5, 7, 11, 13, 17, 19]
//   sievePrimes(1)  -> []
//
// ─────────────────────────────────────────────────────────
// isPrime — CARA KERJA:
//   Naif: cek semua pembagi dari 2 sampai n-1 -> O(n)
//   Optimasi 1: hanya perlu cek sampai sqrt(n) -> O(sqrt n)
//     Kenapa? Kalau n = a*b dan a > sqrt(n), maka b < sqrt(n)
//     Jadi pasti sudah ketemu factor b sebelumnya.
//   Optimasi 2: skip angka genap (kecuali 2) -> loop hanya angka ganjil
//
//   isPrime(49):
//   - 49 bukan 0,1 -> lanjut
//   - 49 % 2 != 0  -> lanjut
//   - sqrt(49) = 7, loop i=3,5,7
//   - i=7: 49 % 7 == 0 -> return false ✓
//
// ─────────────────────────────────────────────────────────
// sievePrimes — CARA KERJA (Sieve of Eratosthenes):
//   1. Buat array boolean ukuran n+1, semua = true (asumsi prima)
//   2. Set index 0 dan 1 = false
//   3. Untuk tiap i dari 2 sampai sqrt(n):
//      - Kalau sieve[i] masih true (prima), tandai semua kelipatan i = false
//      - Mulai dari i*i (bukan 2*i) karena lebih kecil sudah ditandai
//   4. Kumpulkan semua index yang masih true
//
//   sievePrimes(10):
//   i=2: tandai 4,6,8,10 = false
//   i=3: tandai 9 = false (6 sudah ditandai)
//   Hasil: index 2,3,5,7 masih true -> [2,3,5,7]
//
// KUNCI JAWABAN saat ditanya interviewer:
//   isPrime: "Loop sampai sqrt(n), skip even. O(sqrt n) time."
//   sieve: "Mark composite numbers starting from i*i. O(n log log n) time."
//
// EDGE CASES yang harus disebut:
//   - n < 2     -> false (bukan prima)
//   - n = 2     -> true  (prima terkecil, satu-satunya prima genap)
//   - n genap   -> langsung false (kecuali 2)
// =============================================================================

// isPrime: cek satu angka — O(sqrt n) time, O(1) space
func isPrime(n int) bool {
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

// sievePrimes: semua prima sampai n — O(n log log n) time, O(n) space
func sievePrimes(n int) []int {
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
