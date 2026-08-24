package reverseinteger

import "math"

// =============================================================================
// REVERSE INTEGER
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Matematika, digit manipulation, overflow check
// Complexity: O(log x) time — jumlah iterasi = jumlah digit
// =============================================================================
//
// SOAL:
//   Balik digit dari sebuah integer 32-bit.
//   Return 0 jika hasil reverse melebihi range int32. (LeetCode #7)
//
// CONTOH:
//   reverse(123)         -> 321
//   reverse(-123)        -> -321
//   reverse(120)         -> 21    (trailing zero hilang)
//   reverse(0)           -> 0
//   reverse(1534236469)  -> 0     (9646324351 > MaxInt32 = 2147483647)
//   reverse(-1534236469) -> 0     (-9646324351 < MinInt32 = -2147483648)
//
// RANGE INT32:
//   MaxInt32 =  2,147,483,647
//   MinInt32 = -2,147,483,648
//
// EDGE CASES yang harus disebut:
//   - Angka negatif      -> tanda minus ikut (Go % mempertahankan tanda)
//   - Trailing zero      -> hilang secara alami (120 % 10 = 0, 12 % 10 = 2, 1 % 10 = 1)
//   - Overflow           -> return 0
//   - x = 0             -> return 0
//
// CARA KERJA:
//   x=123:
//   iter 1: digit=123%10=3, x=12, result=0*10+3=3
//   iter 2: digit=12%10=2,  x=1,  result=3*10+2=32
//   iter 3: digit=1%10=1,   x=0,  result=32*10+1=321
//   x=0 -> loop selesai
//   321 <= MaxInt32 -> return 321 ✓
//
//   x=-123:
//   iter 1: digit=-123%10=-3, x=-12, result=-3
//   iter 2: digit=-12%10=-2,  x=-1,  result=-3*10+(-2)=-32
//   iter 3: digit=-1%10=-1,   x=0,   result=-32*10+(-1)=-321
//   -321 >= MinInt32 -> return -321 ✓
//
// KENAPA Go % AMAN UNTUK NEGATIF?
//   Di Go, -123 % 10 = -3  (bukan 7 seperti Python)
//   Tanda hasil % mengikuti dividend (angka kiri)
//   Jadi negatif otomatis handled dengan benar
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Ambil digit terakhir dengan x%10, buang dengan x/10.
//    Bangun result: result = result*10 + digit.
//    Di akhir cek overflow: result harus dalam range MaxInt32/MinInt32."
//
// JEBAKAN UMUM:
//   - Cek overflow sebelum result dibangun sepenuhnya (tidak akurat)
//   - Lupa bahwa Go int bisa 64-bit, jadi overflow int32 tidak otomatis terdeteksi
//   - Pakai strconv.Itoa + Reverse + Atoi -> lebih panjang dan perlu handle '-'
// =============================================================================

// reverse: digit reversal — O(log x) time, O(1) space
func reverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10  // ambil digit terakhir (negatif kalau x negatif)
		x /= 10           // buang digit terakhir
		result = result*10 + digit
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
