package Codity
// 100%

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	m := make(map[int]bool, 0)

	for i,  a := range A {
		if !m[a-1] {
			m[a-1] = true
		}
		if len(m) == X {
			return i
		}
	}
	return -1
}