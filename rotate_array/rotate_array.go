package rotatearray

// Rotate Array
// ------------
// Putar array ke kanan sebanyak k langkah.
// (LeetCode #189)
//
// Contoh:
//
//	rotate([]int{1,2,3,4,5,6,7}, 3) -> [5,6,7,1,2,3,4]
//	rotate([]int{-1,-100,3,99}, 2)  -> [3,99,-1,-100]
//
// Bonus: Selesaikan dengan O(1) extra space (3-reverse trick).
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, lo, hi int) {
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
}
