package validparentheses

// =============================================================================
// VALID PARENTHESES
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Stack
// Complexity: O(n) time, O(n) space
// =============================================================================
//
// SOAL:
//   Diberikan string berisi karakter '(', ')', '{', '}', '[', ']'.
//   Cek apakah string tersebut valid (semua bracket tutup dengan benar).
//   (LeetCode #20)
//
// ATURAN VALID:
//   1. Setiap bracket buka harus ditutup dengan bracket yang sama
//   2. Bracket harus ditutup dalam urutan yang benar (LIFO)
//
// CONTOH:
//   isValid("()")        -> true
//   isValid("()[]{}")    -> true
//   isValid("([])")      -> true
//   isValid("(]")        -> false  (bracket tidak cocok)
//   isValid("([)]")      -> false  (urutan tidak benar)
//   isValid("{[]}")      -> true
//   isValid("")          -> true   (string kosong = valid)
//   isValid("(")         -> false  (tidak punya pasangan tutup)
//   isValid(")")         -> false  (langsung closing, stack kosong)
//
// EDGE CASES yang harus disebut:
//   - String kosong                    -> true
//   - Hanya opening bracket "((("      -> false (stack tidak kosong di akhir)
//   - Hanya closing bracket ")))"      -> false (stack kosong saat pop)
//   - Panjang ganjil                   -> langsung false (optimasi)
//
// CARA KERJA (step by step):
//   Input: "([)]"
//   i=0: '(' -> push ke stack. stack=['(']
//   i=1: '[' -> push ke stack. stack=['(','[']
//   i=2: ')' -> closing, cek top stack: '[' != '(' -> FALSE
//
//   Input: "([])"
//   i=0: '(' -> push. stack=['(']
//   i=1: '[' -> push. stack=['(','[']
//   i=2: ']' -> closing, top='[', match -> pop. stack=['(']
//   i=3: ')' -> closing, top='(', match -> pop. stack=[]
//   Stack kosong -> TRUE
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Pakai stack. Kalau ketemu bracket buka, push.
//    Kalau ketemu bracket tutup, cek top stack: harus pasangannya.
//    Kalau tidak cocok atau stack kosong -> false.
//    Di akhir, stack harus kosong -> semua bracket sudah dipasangkan."
//
// JEBAKAN UMUM:
//   - Lupa cek stack kosong sebelum pop -> runtime panic
//   - Lupa cek len(stack)==0 di akhir -> "((" dianggap valid
// =============================================================================

func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		}
	}
	return len(stack) == 0
}
