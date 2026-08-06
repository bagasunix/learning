package majorityelement

// Majority Element
// ----------------
// Cari elemen yang muncul lebih dari n/2 kali dalam array.
// Asumsi: selalu ada majority element. (LeetCode #169)
//
// Contoh:
//
//	majorityElement([]int{3,2,3})           -> 3
//	majorityElement([]int{2,2,1,1,1,2,2})   -> 2
//
// Bonus: Selesaikan dalam O(n) time dan O(1) space (Boyer-Moore Voting).
func majorityElement(nums []int) int {
	candidate, count := nums[0], 1
	for _, v := range nums[1:] {
		if count == 0 {
			candidate = v
			count = 1
		} else if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
