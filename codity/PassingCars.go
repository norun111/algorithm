package codity

// timeout 50%

func Solution(A []int) int {

	count := 0
	for i:=0; i<len(A); i++ {

		for j:=i+1;j<len(A); j++ {
			// fmt.Println(A[i], A[j])
			if A[i] == 0 {
				if A[j] == 1 {
					count++
				}
			}
		}
	}

	if count > 1000000000 {
		return -1
	} else {
		return count
	}
}