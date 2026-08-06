package maxsubarray

// Maximum Subarray (Kadane's Algorithm)
// -------------------------------------
// Cari jumlah maksimum dari contiguous subarray.
// (LeetCode #53)
//
// Contoh:
//
//	maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}) -> 6  ([4,-1,2,1])
//	maxSubArray([]int{1})                     -> 1
//	maxSubArray([]int{5,4,-1,7,8})            -> 23
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	cur := nums[0]
	for _, v := range nums[1:] {
		if cur < 0 {
			cur = v
		} else {
			cur += v
		}
		if cur > maxSum {
			maxSum = cur
		}
	}
	return maxSum
}
