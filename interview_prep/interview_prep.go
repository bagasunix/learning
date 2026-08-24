package interviewprep

// =============================================================================
// INTERVIEW PREP — SOAL PRIORITAS TINGGI
// =============================================================================
// File ini menggabungkan semua soal yang paling mungkin keluar.
// Pelajari urutan dari atas ke bawah.
//
// QUICK REFERENCE — hafal ini sebelum test:
//
//  No | Soal              | Approach              | Time   | Space
//  ---|-------------------|-----------------------|--------|-------
//  1  | FizzBuzz          | loop + modulo         | O(n)   | O(n)
//  2  | Fibonacci         | iterative dua var     | O(n)   | O(1)
//  3  | Palindrome        | two-pointer + clean   | O(n)   | O(n)
//  4  | Anagram           | freq counter          | O(n)   | O(1)
//  5  | Reverse Words     | Fields + two-pointer  | O(n)   | O(n)
//  6  | Two Sum           | hashmap complement    | O(n)   | O(n)
//  7  | Max Subarray      | Kadane's              | O(n)   | O(1)
//  8  | Rotate Array      | 3-reverse trick       | O(n)   | O(1)
//  9  | Valid Parentheses | stack match           | O(n)   | O(n)
//  10 | Binary Search     | lo/hi mid             | O(logn)| O(1)
//  11 | Prime Number      | sqrt loop skip even   | O(√n)  | O(1)
//  12 | Reverse Integer   | %10 digit + overflow  | O(logn)| O(1)
//  13 | Factorial         | iterative             | O(n)   | O(1)
// =============================================================================

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// =============================================================================
// 1. FIZZBUZZ
// =============================================================================
// SOAL: Angka 1 sampai n. Kelipatan 3="Fizz", kelipatan 5="Buzz",
//       kelipatan keduanya="FizzBuzz", selain itu angkanya.
//
// CONTOH:
//   fizzbuzz(5) -> ["1","2","Fizz","4","Buzz"]
//
// KUNCI: Cek 15 dulu, baru 3, baru 5. Kalau terbalik FizzBuzz tidak pernah keluar.
// EDGE CASES: n=0 -> [], n<0 -> []
// =============================================================================

func FizzBuzz(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// =============================================================================
// 2. FIBONACCI
// =============================================================================
// SOAL (variasi 1): Return slice n angka pertama deret fibonacci.
//   fibonacci(5) -> [0,1,1,2,3]
//
// SOAL (variasi 2): Return nilai fibonacci ke-n.
//   fib(6) -> 8   (deret: 0,1,1,2,3,5,8)
//
// KUNCI: Simpan dua variabel a,b saja. a=prev, b=curr. Swap: a,b = b, a+b
// EDGE CASES: n=0 -> 0, n=1 -> 1, n<0 -> tanyakan interviewer
//
// JANGAN pakai rekursi biasa: O(2^n) -> TLE untuk n>40
// =============================================================================

// Fibonacci: return slice n angka pertama
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

// Fib: return nilai ke-n — O(n) time, O(1) space
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// =============================================================================
// 3. PALINDROME
// =============================================================================
// SOAL: Cek apakah string adalah palindrome.
//       Abaikan case dan karakter non-alphanumeric. (LeetCode #125)
//
// CONTOH:
//   IsPalindrome("A man a plan a canal Panama") -> true
//   IsPalindrome("race a car")                  -> false
//   IsPalindrome("")                            -> true
//
// KUNCI: Clean dulu (lowercase + alphanumeric only), lalu two-pointer.
// EDGE CASES: string kosong -> true, satu karakter -> true
// =============================================================================

func IsPalindrome(s string) bool {
	// step 1: clean — lowercase, buang non-alphanumeric
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	cleaned := b.String()

	// step 2: two-pointer dari luar ke dalam
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

// =============================================================================
// 4. ANAGRAM
// =============================================================================
// SOAL: Cek apakah dua string adalah anagram. (LeetCode #242)
//
// CONTOH:
//   IsAnagram("listen", "silent") -> true
//   IsAnagram("hello", "world")   -> false
//   IsAnagram("", "")             -> true
//
// KUNCI: Frequency counter — increment untuk A, decrement untuk B.
//        Kalau ada counter negatif -> bukan anagram.
// EDGE CASES: panjang beda -> langsung false
// =============================================================================

func IsAnagram(a, b string) bool {
	// early exit: panjang beda pasti bukan anagram
	if len(a) != len(b) {
		return false
	}
	counter := make(map[rune]int)
	for _, c := range strings.ToLower(a) {
		counter[c]++
	}
	for _, c := range strings.ToLower(b) {
		counter[c]--
		if counter[c] < 0 {
			return false
		}
	}
	return true
}

// =============================================================================
// 5. REVERSE WORDS
// =============================================================================
// SOAL: Balik urutan kata. Hapus spasi berlebih. (LeetCode #151)
//
// CONTOH:
//   ReverseWords("the sky is blue")  -> "blue is sky the"
//   ReverseWords("  hello world  ")  -> "world hello"
//   ReverseWords("a good   example") -> "example good a"
//
// KUNCI: strings.Fields handle multiple spaces otomatis.
//        Two-pointer untuk reverse slice kata.
// EDGE CASES: semua spasi -> "", satu kata -> kata itu
// =============================================================================

func ReverseWords(s string) string {
	words := strings.Fields(s) // split + strip whitespace otomatis
	lo, hi := 0, len(words)-1
	for lo < hi {
		words[lo], words[hi] = words[hi], words[lo]
		lo++
		hi--
	}
	return strings.Join(words, " ")
}

// =============================================================================
// 6. TWO SUM
// =============================================================================
// SOAL: Cari dua index i,j sehingga nums[i]+nums[j]==target. (LeetCode #1)
//
// CONTOH:
//   TwoSum([2,7,11,15], 9)  -> [0,1]  (2+7=9)
//   TwoSum([3,2,4], 6)      -> [1,2]  (2+4=6)
//   TwoSum([3,3], 6)        -> [0,1]
//
// KUNCI: hashmap {nilai -> index}. Untuk tiap angka, cari complement = target-angka.
//        Kalau complement ada di map -> return. Kalau tidak -> simpan ke map.
// EDGE CASES: dua elemen sama ([3,3],target=6) -> [0,1]
//
// JANGAN pakai nested loop O(n²) — interviewer harapkan O(n).
// =============================================================================

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index
	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

// =============================================================================
// 7. MAX SUBARRAY (Kadane's Algorithm)
// =============================================================================
// SOAL: Cari jumlah maksimum subarray berurutan. (LeetCode #53)
//
// CONTOH:
//   MaxSubArray([-2,1,-3,4,-1,2,1,-5,4]) -> 6  ([4,-1,2,1])
//   MaxSubArray([1])                     -> 1
//   MaxSubArray([-1,-2,-3])             -> -1  (semua negatif, ambil terbesar)
//
// KUNCI: kalau cur negatif, mulai fresh dari elemen berikutnya.
//        cur = max(v, cur+v). Track maxSum global.
// EDGE CASES: semua negatif -> return angka paling besar (kurang negatif)
//             inisialisasi maxSum = nums[0], BUKAN 0!
// =============================================================================

func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	cur := nums[0]
	for _, v := range nums[1:] {
		if cur < 0 {
			cur = v // mulai fresh
		} else {
			cur += v
		}
		if cur > maxSum {
			maxSum = cur
		}
	}
	return maxSum
}

// =============================================================================
// 8. ROTATE ARRAY
// =============================================================================
// SOAL: Putar array ke kanan sebanyak k langkah. (LeetCode #189)
//
// CONTOH:
//   Rotate([1,2,3,4,5,6,7], 3) -> [5,6,7,1,2,3,4]
//   Rotate([-1,-100,3,99], 2)  -> [3,99,-1,-100]
//
// KUNCI 3-reverse trick:
//   1. Reverse semua         [7,6,5,4,3,2,1]
//   2. Reverse [0..k-1]      [5,6,7,4,3,2,1]
//   3. Reverse [k..n-1]      [5,6,7,1,2,3,4] ✓
// EDGE CASES: k >= n -> k = k%n, k=0 -> tidak berubah
// =============================================================================

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

// =============================================================================
// 9. VALID PARENTHESES
// =============================================================================
// SOAL: Cek apakah string bracket valid (semua tertutup dengan benar). (LeetCode #20)
//
// CONTOH:
//   IsValid("()")        -> true
//   IsValid("()[]{}")    -> true
//   IsValid("(]")        -> false
//   IsValid("([)]")      -> false
//   IsValid("")          -> true
//
// KUNCI: push opening bracket. Saat closing, pop dan cek cocok tidak.
//        Di akhir stack harus kosong.
// EDGE CASES: stack kosong saat mau pop -> false
//             stack tidak kosong di akhir -> false
// =============================================================================

func IsValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		}
	}
	return len(stack) == 0
}

// =============================================================================
// 10. BINARY SEARCH
// =============================================================================
// SOAL: Cari index target dalam sorted array. Return -1 jika tidak ada. (LeetCode #704)
//
// CONTOH:
//   BinarySearch([1,3,5,7,9,11], 7)  -> 3
//   BinarySearch([1,3,5,7,9,11], 4)  -> -1
//
// KUNCI: mid = lo + (hi-lo)/2  <- WAJIB, bukan (lo+hi)/2 karena overflow!
//        lo<=hi (bukan lo<hi, agar satu elemen tidak kelewat).
// EDGE CASES: array kosong -> -1, satu elemen -> cek langsung
// =============================================================================

func BinarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // aman dari overflow
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

// =============================================================================
// 11. PRIME NUMBER
// =============================================================================
// SOAL: Cek apakah n adalah bilangan prima.
//
// CONTOH:
//   IsPrime(2)  -> true
//   IsPrime(7)  -> true
//   IsPrime(1)  -> false
//   IsPrime(9)  -> false  (9 = 3*3)
//
// KUNCI: loop hanya sampai sqrt(n), skip angka genap.
// EDGE CASES: n<2 -> false, n=2 -> true (satu-satunya prima genap)
// =============================================================================

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
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// =============================================================================
// 12. REVERSE INTEGER
// =============================================================================
// SOAL: Balik digit integer 32-bit. Return 0 jika overflow. (LeetCode #7)
//
// CONTOH:
//   ReverseInt(123)         -> 321
//   ReverseInt(-123)        -> -321
//   ReverseInt(120)         -> 21
//   ReverseInt(1534236469)  -> 0  (overflow)
//
// KUNCI: digit = x%10, x /= 10, result = result*10 + digit.
//        Go % mempertahankan tanda negatif, jadi negatif otomatis handled.
// EDGE CASES: overflow MaxInt32/MinInt32 -> return 0
// =============================================================================

func ReverseInt(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		result = result*10 + digit
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

// =============================================================================
// 13. FACTORIAL
// =============================================================================
// SOAL: Hitung n! = n * (n-1) * ... * 2 * 1
//
// CONTOH:
//   Factorial(0)  -> 1   (by definition)
//   Factorial(5)  -> 120
//   Factorial(10) -> 3628800
//
// KUNCI: iterative lebih baik dari rekursif (tidak ada stack overhead).
//        0! = 1 by definition matematika.
// EDGE CASES: n=0 -> 1, n=1 -> 1, n>20 -> overflow int64
// =============================================================================

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// =============================================================================
// BONUS: PRINT CHEAT SHEET
// =============================================================================
// Jalankan fungsi ini untuk lihat ringkasan di terminal.
// =============================================================================

func PrintCheatSheet() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          INTERVIEW PREP — CHEAT SHEET                   ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1  FizzBuzz       : cek %15 dulu, baru %3, baru %5      ║")
	fmt.Println("║ 2  Fibonacci      : a,b = b, a+b  O(n) O(1)             ║")
	fmt.Println("║ 3  Palindrome     : clean -> two-pointer lo/hi           ║")
	fmt.Println("║ 4  Anagram        : freq map +1 A, -1 B, negatif=false  ║")
	fmt.Println("║ 5  Reverse Words  : Fields() + reverse slice + Join()   ║")
	fmt.Println("║ 6  Two Sum        : hashmap complement = target-num      ║")
	fmt.Println("║ 7  Max Subarray   : Kadane: cur<0 reset, track max       ║")
	fmt.Println("║ 8  Rotate Array   : k%n, reverse all, rev[0:k], rev[k:] ║")
	fmt.Println("║ 9  Valid Parens   : push open, pop&match close, len==0  ║")
	fmt.Println("║ 10 Binary Search  : mid=lo+(hi-lo)/2, lo<=hi            ║")
	fmt.Println("║ 11 Prime Number   : loop sqrt(n), skip even             ║")
	fmt.Println("║ 12 Reverse Int    : %10 digit, /10 strip, check overflow ║")
	fmt.Println("║ 13 Factorial      : loop i=2..n, result*=i              ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
}
