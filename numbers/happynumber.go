// Write an algorithm to determine if a number n is happy.
// A happy number is a number defined by the following process:
// Starting with any positive integer, replace the number by the sum of the
// squares of its digits. Repeat the process until the number equals 1
// (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

package numbers

var visited map[int]struct{}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	visited = map[int]struct{}{}
	for {
		visited[n] = struct{}{}
		if n == 1 {
			break
		}
		n = sum(n)
		if _, ok := visited[n]; ok {
			return false
		}
	}
	return true
}
func sum(n int) int {
	s := 0
	for {
		if n == 0 {
			break
		}
		s += (n % 10) * (n % 10)
		n = n / 10
	}
	return s
}
