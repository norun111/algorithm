package Codity

func Solution(A []int) int {
	// write your code in Go 1.4
	should_be := len(A)
	sum_is := 0

	for idx, _ := range A {
		sum_is += A[idx]
		should_be += idx+1
	}

	return should_be - sum_is +1
}