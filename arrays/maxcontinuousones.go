// Given a binary array nums, return the maximum number of consecutive 1's
// in the array.
package arrays

func findMaxConsecutiveOnes(nums []int) int {
	start := 0
	end := 0
	max := 0
	for i := range nums {
		if nums[i] == 0 {
			start++
			end = start
		} else {
			if end-start+1 > max {
				max = end - start + 1
			}
			end++

		}
	}
	return max
}
