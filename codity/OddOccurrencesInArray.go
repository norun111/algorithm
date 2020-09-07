package Codity


// 100%
import (
	"sort"
)

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)

	for i:=0; i<len(A)-1; {
		if A[i] != A[i+1] {

			return A[i]
		}
		i = i+2
	}
	return A[len(A)-1]
}

// Timeout

func Solution(A []int) int {
	// write your code in Go 1.4
	m := make(map[int]int, 0)
	for _, a := range A {
		m[a] = 0
	}


	for key, _ := range m {
		for _, a := range A {
			if key == a {
				m[key]++
			}
		}
	}

	result := 0
	for key, _ := range m {
		if m[key]%2!=0 {
			result = key
		}
	}

	return result
}