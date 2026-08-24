package fizzbuzz

import "strconv"

// =============================================================================
// FIZZBUZZ
// =============================================================================
// Tingkat   : Mudah
// Konsep    : Loop, kondisional, modulo (%)
// Complexity: O(n) time, O(n) space
// =============================================================================
//
// SOAL:
//   Tulis fungsi yang menerima integer n, dan return slice string
//   berisi angka 1 sampai n dengan aturan:
//     - Kelipatan 3        -> "Fizz"
//     - Kelipatan 5        -> "Buzz"
//     - Kelipatan 3 DAN 5  -> "FizzBuzz"
//     - Selain itu         -> angkanya sebagai string
//
// CONTOH:
//   fizzbuzz(5)  -> ["1", "2", "Fizz", "4", "Buzz"]
//   fizzbuzz(15) -> ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz",
//                    "11","Fizz","13","14","FizzBuzz"]
//
// EDGE CASES yang harus disebut ke interviewer:
//   - n = 0  -> return slice kosong []
//   - n = 1  -> return ["1"]
//   - n < 0  -> bisa return [] atau error, tergantung kesepakatan
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Cek kelipatan 15 dulu (3*5), baru 3, baru 5.
//    Kalau 3 dan 5 dicek duluan, FizzBuzz tidak akan pernah tercapai."
//
// JEBAKAN UMUM:
//   - Cek i%3 dan i%5 secara terpisah pakai if-else if -> FizzBuzz ketinggalan
//   - Urutan kondisi salah (15 harus duluan)
// =============================================================================

func fizzbuzz(n int) []string {
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
