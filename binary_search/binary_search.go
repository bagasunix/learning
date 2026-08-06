package binarysearch

// Binary Search
// --------------
// Cari index dari target dalam sorted array.
// Kembalikan -1 kalau tidak ditemukan.
//
// Contoh:
//
//	binarySearch([]int{1,3,5,7,9,11}, 7)  -> 3
//	binarySearch([]int{1,3,5,7,9,11}, 4)  -> -1
//
// Bonus: Implementasikan versi recursive.
func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
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

// BinarySearchRecursive versi recursive.
func binarySearchRecursive(nums []int, target int) int {
	return bsHelper(nums, target, 0, len(nums)-1)
}

func bsHelper(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return bsHelper(nums, target, mid+1, hi)
	}
	return bsHelper(nums, target, lo, mid-1)
}
