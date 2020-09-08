package Codity

// performance is OK

import (
	// "fmt"
	"sort"
)


func removeDuplicate(args []int) []int {
	results := make([]int, 0, len(args))
	encountered := map[int]bool{}
	for i := 0; i < len(args); i++ {
		if args[i] > 0 {
			if !encountered[args[i]] {
				encountered[args[i]] = true
				results = append(results, args[i])
			}
		}

	}
	return results
}

func Solution(A []int) int {
	sort.Ints(A)

	if A[len(A)-1] < 0 {
		return 1
	}
	A = removeDuplicate(A)


	condition := false

	for i:=0; i<len(A)-1; i++{
		if A[i]+1 != A[i+1] {
			condition = true
			return A[i]+1
		}
	}
	if !condition {
		return A[len(A)-1]+1
	}
	return -1
}