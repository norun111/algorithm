package codity

func Solution(A []int) int {
	// write your code in Go 1.4
	m := make(map[int]int, 0)

	if len(A) == 1 {
		return A[0]
	}

	for i:=0; i<len(A); i++ {
		m[A[i]] = 0
	}

	for key, _ := range m {
		for _, a := range A {
			if key == a {
				m[key]++
			}
		}
	}
	result := 0

	for key, value := range m {
		if value%2 != 0 {
			result = key
		}
	}

	return result
}