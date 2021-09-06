// Given two integer arrays nums1 and nums2, return an array of their
// intersection. Each element in the result must be unique and you may return
// the result in any order.
package arrays

var ret []int

func common(nums []int, m map[int]bool) {
	for _, val := range nums {
		if vl, ok := m[val]; ok {
			if !vl {
				ret = append(ret, val)
				m[val] = true
			}
		}
	}
}

func Intersection(nums1 []int, nums2 []int) []int {
	ret = []int{}
	m := map[int]bool{}
	l1 := len(nums1)
	l2 := len(nums2)

	if l1 < l2 {
		for i := 0; i < l1; i++ {
			m[nums1[i]] = false
		}
		common(nums2, m)
	} else {
		for i := 0; i < l2; i++ {
			m[nums2[i]] = false
		}
		common(nums1, m)
	}
	return ret
}
