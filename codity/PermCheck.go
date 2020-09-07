package Codity

// 88%

// you can also use imports, for example:
import (
	"sort"
	// "fmt"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)

	result := 1

	if len(A) == 1 {
		if A[0] == 1 {
			return 1
		} else {
			return 0
		}
	}

	for i, _ := range A {
		if i > 0 {
			if A[i] != A[i-1]+1 {
				result = 0
				break
			}
		}
	}

	return result
}