// Given an integer array nums and an integer k, return true if there are two
// distinct indices i and j in the array such that nums[i] == nums[j] and
// abs(i - j) <= k.
package arrays

func containsNearbyDuplicate(nums []int, k int) bool {
	start := 0

	for start < len(nums) {
		end := start + 1
		for end <= start+k {
			if end >= len(nums) {
				break
			}
			if nums[start] == nums[end] {
				return true
			}
			end++
		}
		start++
	}
	return false
}
