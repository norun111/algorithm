package Codity

// Timeout Error

import (
	"strings"
	"sort"
)

func Solution(S string, P []int, Q []int) []int {
	// write your code in Go 1.4
	strSlice := strings.Split(S, "")

	result := []int{}

	for i:=0; i<len(P); i++ {
		first := P[i]
		second := Q[i]

		minList := make([]int, 0)
		for j:=first; j<=second; j++ {
			if strSlice[j] == "A" {
				minList = append(minList, 1)
			} else if strSlice[j] == "C" {
				minList = append(minList, 2)
			} else if strSlice[j] == "G" {
				minList = append(minList, 3)
			} else if strSlice[j] == "T" {
				minList = append(minList, 4)
			}
		}

		sort.Ints(minList)

		result = append(result, minList[0])
	}

	return result
}