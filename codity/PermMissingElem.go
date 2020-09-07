package codity

// you can also use imports, for example:
import (
	"sort"
	// "fmt"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	sort.Ints(A)

	nowNum := 0


	for i:=0; i<len(A); i++ {
		if i==0 {
			nowNum = A[i]
		} else {
			if A[i] == nowNum+1 {
				nowNum = A[i]
			} else {
				nowNum = nowNum + 1
				break
			}
		}
	}

	return nowNum
}