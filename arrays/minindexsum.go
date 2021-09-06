// Suppose Andy and Doris want to choose a restaurant for dinner, and they both
// have a list of favorite restaurants represented by strings.
// You need to help them find out their common interest with the least list
// index sum. If there is a choice tie between answers, output all of them with
// no order requirement. You could assume there always exists an answer.
package arrays

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	least := []string{}
	candidates := []string{}
	l := len(list1)
	if len(list2) > len(list1) {
		l = len(list2)
	}
	min := 2 * l
	for i, val := range list1 {
		m[val] = i
	}
	for i, val := range list2 {
		if _, ok := m[val]; ok {
			m[val] += i
			if m[val] <= min {
				min = m[val]
				candidates = append(candidates, val)
			}

		}
	}
	for _, v := range candidates {
		if m[v] <= min {
			least = append(least, v)
			min = m[v]
		}
	}
	return least
}
