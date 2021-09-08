// Given an integer array nums and an integer k, return the k most frequent
// elements. You may return the answer in any order.
package arrays

import "sort"

type elem struct {
	val   int
	count int
}
type elemSorter []elem

func (e elemSorter) Len() int {
	return len(e)
}

func (e elemSorter) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e elemSorter) Less(i, j int) bool {
	return e[i].count > e[j].count
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]elem{}
	for i := range nums {
		if v, ok := m[nums[i]]; ok {
			v.count++
			m[nums[i]] = v
		} else {
			m[nums[i]] = elem{nums[i], 1}
		}
	}
	elems := []elem{}
	for _, v := range m {
		elems = append(elems, v)
	}
	sort.Sort(elemSorter(elems))
	ret := []int{}
	for i := 0; i < k; i++ {
		ret = append(ret, elems[i].val)
	}
	return ret
}
