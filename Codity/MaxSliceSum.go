package Codity

import "sort"

func Solution(A []int) int {
	// write your code in Go 1.4
	max_ending := -1000000
	max_slice := -1000000

	for i:=0; i<len(A); i++ {
		slice1 := []int{A[i], max_ending+A[i]}
		slice2 := []int{max_slice, max_ending}
		sort.Ints(slice1)
		sort.Ints(slice2)
		max_ending = slice1[len(slice1)-1]
		max_slice = slice2[len(slice2)-1]
	}

	return max_slice
}