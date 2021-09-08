// Given an array nums of integers, return how many of them contain an even
// number of digits.
package arrays

func calcDigits(num int) int {
	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	return count
}

func findNumbers(nums []int) int {
	count := 0
	for i := range nums {
		if calcDigits(nums[i])%2 == 0 {
			count++
		}
	}
	return count
}
