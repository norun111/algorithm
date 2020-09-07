package Codity

// 69% runtimeError

import (
	"math"
)

func Solution(A []int) int {

	min := math.MaxFloat64

	for i:=0; i<len(A)-1; i++ {

		first := 0
		second := 0
		for j:=0; j<i+1; j++ {
			first+=A[j]
		}

		for k:=i+1; k<len(A); k++ {
			second+=A[k]
		}

		result := math.Abs(float64(first-second))

		if result < min {
			min = result
		}
	}

	return int(min)
}