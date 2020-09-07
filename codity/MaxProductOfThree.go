package Codity
// 100%

import (
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)

	maxThree := A[len(A)-3:]

	minTwo := A[0:2]

	sum1 := 1
	for _, max := range maxThree {
		sum1*=max
	}

	sum2 := minTwo[0]*minTwo[1]*maxThree[len(maxThree)-1]
	result := 0

	if sum1 > sum2 {
		result = sum1
	} else if sum1 < sum2 {
		result = sum2
	} else if sum1 == sum2 {
		result = sum1
	}

	return result
}