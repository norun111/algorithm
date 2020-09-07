package Codity

func Solution(A []int, B []int) int {

	stack := make([]int, 0)
	survivals := 0

	for i, _ := range A {
		if B[i] == 0 {
			for len(stack) != 0 {
				if stack[len(stack)-1] > A[i] {
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if len(stack) == 0 {
				survivals += 1
			}

		} else {
			stack = append(stack, A[i])
		}
	}
	survivals+=len(stack)

	return survivals
}