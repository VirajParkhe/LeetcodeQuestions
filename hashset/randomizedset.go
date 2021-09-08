// Implement the RandomizedSet class:
// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present.
// Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns
// true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements
// (it's guaranteed that at least one element exists when this method is
// called). Each element must have the same probability of being returned.
// You must implement the functions of the class such that each function works
// in average O(1) time complexity.
package hashset

import "math/rand"

type RandomizedSet struct {
	m    map[int]int
	list []int
	l    int
}

func Constructor() RandomizedSet {
	rand.Seed(100)
	return RandomizedSet{map[int]int{}, []int{}, 0}
}

func (this *RandomizedSet) getRandom() int {
	return this.list[rand.Intn(this.l)]
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.list = append(this.list, val)
		this.l++
		this.m[val] = this.l - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if v, ok := this.m[val]; ok {
		this.list[v] = this.list[this.l-1]
		this.m[this.list[v]] = v
		this.list = this.list[:this.l-1]
		this.l--
		delete(this.m, val)
		return true

	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	val := this.getRandom()
	return val
}
