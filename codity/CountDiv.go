package codity

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	numList := make([]int, 0)
	for i:=A; i<=B; i++ {
		numList = append(numList, i)
	}

	count := 0

	for _, num := range numList {
		if num%K == 0 {
			count++
		}
	}

	return count
}