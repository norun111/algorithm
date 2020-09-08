package Codity

func Solution(A []int) []int {
	result := []int{}

	for i:=0; i<len(A); i++ {

		count := 0
		for j:=0; j<len(A); j++ {

			if A[i]%A[j]!=0 {
				count++
			}
		}

		result = append(result, count)
	}

	return result

}