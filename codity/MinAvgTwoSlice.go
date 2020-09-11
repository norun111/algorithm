package Codity
// Timeout Error 60%

// you can also use imports, for example:
import (
	"math"
	// "fmt"
)

func Solution(A []int) int {
	// write your code in Go 1.4
	maxNum := math.MaxFloat64
	idx := 0

	for i:=0; i<len(A); i++ {

		for j:=i+1; j<len(A); j++ {
			sum := 0
			cnt := 0
			for k:=i; k<=j; k++ {
				cnt++
				sum+=A[k]
			}
			result := float64(sum)/float64(cnt)

			if result < maxNum {
				maxNum = result
				idx = i
			}
		}
	}

	return idx
}