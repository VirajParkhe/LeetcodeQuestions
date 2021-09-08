// You are given a string s of lowercase English letters and an integer array
// shifts of the same length. Call the shift() of a letter, the next letter in
// the alphabet, (wrapping around so that 'z' becomes 'a'). For example,
// shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'. Now for each
// shifts[i] = x, we want to shift the first i + 1 letters of s, x times.
// Return the final string after all such shifts to s are applied.
package strings

func shiftingLetters(s string, shifts []int) string {
	sum := 0
	for i := len(shifts) - 1; i >= 0; i-- {
		temp := shifts[i]
		shifts[i] += sum
		shifts[i] = shifts[i] % 26
		sum += temp
	}
	d := ""
	for i := range shifts {
		total := (shifts[i] + int(s[i]))
		if total > 122 {
			total = total - 123 + 97
		}
		d += string(rune(total))
	}
	return d
}
