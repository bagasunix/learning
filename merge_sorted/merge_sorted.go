package mergesorted

// =============================================================================
// MERGE SORTED ARRAYS
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Two-pointer, merge step (dasar dari Merge Sort)
// Complexity: O(n+m) time, O(n+m) space
// =============================================================================
//
// SOAL:
//   Diberikan dua sorted array, gabungkan menjadi satu sorted array.
//   (LeetCode #88 mirip, versi ini return array baru)
//
// CONTOH:
//   mergeSorted([1,3,5], [2,4,6])    -> [1,2,3,4,5,6]
//   mergeSorted([1,2,3], [])         -> [1,2,3]
//   mergeSorted([], [4,5,6])         -> [4,5,6]
//   mergeSorted([1,1,1], [1,1,1])    -> [1,1,1,1,1,1]
//   mergeSorted([-3,-1,0], [-2,2,4]) -> [-3,-2,-1,0,2,4]
//
// EDGE CASES yang harus disebut:
//   - Salah satu array kosong   -> return yang tidak kosong
//   - Kedua array kosong        -> return []
//   - Array dengan duplikat     -> tetap valid
//   - Satu array semua lebih kecil/besar dari yang lain
//
// CARA KERJA (step by step):
//   A=[1,3,5], B=[2,4,6]
//   i=0,j=0: A[0]=1 < B[0]=2  -> append 1, i=1
//   i=1,j=0: A[1]=3 > B[0]=2  -> append 2, j=1
//   i=1,j=1: A[1]=3 < B[1]=4  -> append 3, i=2
//   i=2,j=1: A[2]=5 > B[1]=4  -> append 4, j=2
//   i=2,j=2: A[2]=5 < B[2]=6  -> append 5, i=3
//   i=3: loop selesai (i>=len(A))
//   Append sisa B: [6]
//   Hasil: [1,2,3,4,5,6] ✓
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Dua pointer i dan j, masing-masing untuk tiap array.
//    Bandingkan A[i] dan B[j], append yang lebih kecil, advance pointernya.
//    Setelah salah satu habis, append semua sisa array lainnya.
//    O(n+m) karena setiap elemen dikunjungi tepat satu kali."
//
// JEBAKAN UMUM:
//   - Lupa append sisa array setelah loop utama selesai
//   - Loop condition pakai && bukan bertahap -> bisa miss elemen
// =============================================================================

// mergeSorted: two-pointer — O(n+m) time, O(n+m) space
func mergeSorted(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	// append sisa — salah satu pasti sudah habis
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}
