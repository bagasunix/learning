package sortarray

// =============================================================================
// SORT ARRAY FROM SCRATCH
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Algoritma sorting manual (tanpa sort.Ints / sort.Slice)
// =============================================================================
//
// SOAL:
//   Urutkan array integer secara ascending TANPA menggunakan sort bawaan Go.
//   Implementasikan minimal dua algoritma sorting dari scratch.
//
// CONTOH:
//   bubbleSort([5,3,1,4,2])    -> [1,2,3,4,5]
//   insertionSort([5,3,1,4,2]) -> [1,2,3,4,5]
//   bubbleSort([])             -> []
//   insertionSort([-3,0,-1])   -> [-3,-1,0]
//
// ═══════════════════════════════════════════════════
// ALGORITMA 1: BUBBLE SORT
// ═══════════════════════════════════════════════════
// Complexity: O(n²) worst & average, O(n) best (sudah sorted), O(1) space
//
// Cara kerja:
//   Bandingkan pasangan adjacent (berdampingan), swap kalau urutan salah.
//   Setelah pass pertama, elemen terbesar "menggelembung" ke posisi paling kanan.
//   Ulangi untuk n-1 pass. Optimasi: stop lebih awal kalau tidak ada swap.
//
//   [5,3,1,4,2]:
//   Pass 1: [3,5,1,4,2] -> [3,1,5,4,2] -> [3,1,4,5,2] -> [3,1,4,2,5]  (5 ke ujung)
//   Pass 2: [1,3,4,2,5] -> [1,3,2,4,5]                                 (4 ke tempatnya)
//   Pass 3: [1,2,3,4,5]  swapped=false -> berhenti lebih awal ✓
//
// ═══════════════════════════════════════════════════
// ALGORITMA 2: INSERTION SORT
// ═══════════════════════════════════════════════════
// Complexity: O(n²) worst, O(n) best (sudah sorted), O(1) space
//
// Cara kerja:
//   Ambil elemen satu per satu dari kiri ke kanan.
//   Sisipkan ke posisi yang tepat di bagian kiri yang sudah sorted.
//   Analogi: seperti mengurutkan kartu di tangan.
//
//   [5,3,1,4,2]:
//   i=1: key=3, geser 5 -> [5,5,1,4,2], sisip 3 -> [3,5,1,4,2]
//   i=2: key=1, geser 5,3 -> sisip 1 -> [1,3,5,4,2]
//   i=3: key=4, geser 5 -> sisip 4 -> [1,3,4,5,2]
//   i=4: key=2, geser 5,4,3 -> sisip 2 -> [1,2,3,4,5] ✓
//
// PERBANDINGAN:
//   Bubble Sort   : banyak swap, kurang efisien dalam praktik
//   Insertion Sort: lebih sedikit operasi, bagus untuk data hampir sorted
//   Kedua: O(n²) worst case, tapi insertion sort umumnya lebih cepat
//
// KUNCI JAWABAN saat ditanya interviewer:
//   Bubble: "Bandingkan adjacent, gelembungkan elemen terbesar. Flag early exit."
//   Insertion: "Ambil key, geser elemen lebih besar ke kanan, sisipkan key."
//
// JEBAKAN UMUM:
//   Bubble:    - Lupa flag `swapped` -> tidak ada early exit optimasi
//   Insertion: - Inner loop condition `j >= 0 && result[j] > key` keduanya wajib
// =============================================================================

// bubbleSort: O(n²) worst, O(n) best — dengan early exit optimization
func bubbleSort(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	n := len(result)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		if !swapped {
			break // sudah sorted, tidak perlu iterasi lanjut
		}
	}
	return result
}

// insertionSort: O(n²) worst, O(n) best — bagus untuk data hampir sorted
func insertionSort(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	n := len(result)
	for i := 1; i < n; i++ {
		key := result[i]
		j := i - 1
		// geser semua elemen yang lebih besar dari key ke kanan satu posisi
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key // sisipkan key di posisi yang tepat
	}
	return result
}
