package Codity

import (
	"math"
	// "fmt"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func unset(s []int, i int) []int {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func Solution(A []int) int {
	nowMax := A[0]
	cnt := 0

	indexList := make([]int, 0)

	for i:=1; i<len(A); i++ {
		if A[i] > nowMax {
			nowMax = A[i]
		} else {
			// A[i-1] is Peak
			// fmt.Println(A[i-1])
			indexList = append(indexList, i-1)
			cnt++
			nowMax = A[i]
		}
	}

	minDif := math.MaxInt64

	for i:=1; i<len(indexList); i++ {
		dif := indexList[i] - indexList[i-1]

		if minDif > dif {
			minDif = dif
		}
	}

	for i:=0; i<len(indexList); i++ {


		for j:=i+1; j<len(indexList); j++ {
			dif := indexList[j] - indexList[i]
			if dif > minDif {

			} else {
				indexList = unset(indexList, j)
			}
		}
	}

	return len(indexList)
}