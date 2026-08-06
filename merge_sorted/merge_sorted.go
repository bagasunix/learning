package mergesorted

// Merge Sorted Arrays
// -------------------
// Gabungkan dua sorted array menjadi satu sorted array.
//
// Contoh:
//
//	mergeSorted([]int{1,3,5}, []int{2,4,6}) -> [1,2,3,4,5,6]
//	mergeSorted([]int{1,2,3}, []int{})      -> [1,2,3]
//
// Hint: Two-pointer technique, O(n+m).
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
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}
