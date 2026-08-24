package twosum

// =============================================================================
// TWO SUM
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Hashmap, nested loop (brute force)
// =============================================================================
//
// SOAL:
//   Diberikan array integer dan sebuah target, cari dua index i dan j
//   sehingga nums[i] + nums[j] == target. Dijamin ada tepat satu solusi.
//   (LeetCode #1)
//
// CONTOH:
//   twoSum([2,7,11,15], 9)  -> [0,1]  (2+7=9)
//   twoSum([3,2,4], 6)      -> [1,2]  (2+4=6)
//   twoSum([3,3], 6)        -> [0,1]  (3+3=6)
//
// EDGE CASES yang harus disebut:
//   - Dua elemen sama (nums=[3,3], target=6)  -> [0,1]
//   - Target lebih besar dari semua angka     -> tidak ada (soal bilang pasti ada)
//   - Array panjang 2                         -> selalu index [0,1]
//
// ADA DUA IMPLEMENTASI:
//
//   1. twoSum (brute force):
//      - Double loop, cek semua pasangan
//      - O(n²) time, O(1) space
//      - Mulai dari sini saat interview, lalu tawarkan optimasi
//
//   2. twoSumHashmap (optimal):
//      - Satu loop: untuk tiap nums[i], cek apakah (target - nums[i]) sudah ada di map
//      - Kalau ada -> ketemu, return indexnya
//      - Kalau tidak -> simpan nums[i] ke map dengan valuenya = index i
//      - O(n) time, O(n) space
//      - INI YANG DIHARAPKAN interviewer
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Brute force O(n²): cek semua pasangan.
//    Optimasi: pakai hashmap — untuk tiap angka, hitung complement = target - angka.
//    Kalau complement sudah ada di map, return kedua indexnya.
//    Kalau belum, simpan angka sekarang ke map. Single pass O(n)."
//
// JEBAKAN UMUM:
//   - Lupa bahwa tidak boleh pakai index yang sama dua kali (i != j)
//   - Return nilai angkanya, bukan indexnya
// =============================================================================

// twoSum: brute force — O(n²) time, O(1) space
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// twoSumHashmap: optimal — O(n) time, O(n) space
func twoSumHashmap(nums []int, target int) []int {
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
