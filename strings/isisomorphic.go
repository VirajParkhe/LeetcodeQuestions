// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to
// get t. All occurrences of a character must be replaced with another
// character while preserving the order of characters.
// No two characters may map to the same character, but a character may map to
// itself.
package strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]rune{}
	m2 := map[rune]rune{}
	for i, val := range s {
		if v, ok := m[val]; ok {
			if v != rune(t[i]) {
				return false
			}
		} else {
			if _, ok := m2[rune(t[i])]; !ok {
				m[val] = rune(t[i])
				m2[rune(t[i])] = val
			} else {
				return false
			}
		}
	}
	return true
}
