// You're given strings jewels representing the types of stones that are jewels,
// and stones representing the stones you have. Each character in stones is a
// type of stone you have. You want to know how many of the stones you have are
// also jewels. Letters are case sensitive, so "a" is considered a different
// type of stone from "A".
package strings

func numJewelsInStones(jewels string, stones string) int {
	m := map[rune]int{}
	for _, s := range stones {
		m[s]++
	}
	count := 0
	for _, s := range jewels {
		count += m[s]
	}
	return count
}
