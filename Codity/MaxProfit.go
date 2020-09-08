package Codity

// timeout error

func Solution(A []int) int {
	// write your code in Go 1.4
	maxProf := 0

	for i:=0; i<len(A); i++ {

		for j:=i+1; j<len(A); j++ {

			prof := A[j] - A[i]
			if prof > maxProf {
				maxProf = prof
			}
		}
	}

	if maxProf == 0 {
		return 0
	}
	return maxProf
}