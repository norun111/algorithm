package Codity

// performance is OK

import (
	// "fmt"
	"sort"
	"math"
)

// 66%

func isContain(slice []int, str int) bool {

	condition := false
	for _, s := range slice {

		if s == str {
			condition = true
			break //一番初めだけ検索する
		}
	}

	return condition //なければfalseを返す
}

func Solution(A []int) int {
	// write your code in Go 1.4
	maxValue := math.MaxInt64

	for i:=1; i<maxValue; i++ {
		if !isContain(A, i) {
			return i
		}
	}
	return -1

}

// 55%

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