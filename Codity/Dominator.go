package Codity

// my answer 33%
import (
	// "fmt"
	"sort"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func isDif(slice []int) bool {
	in := 0
	for i, s := range slice {
		if i==0 {
			in = s
		} else {
			if in != s {
				return true
			}
		}
	}
	return false
}

func Solution(A []int) int {
	// write your code in Go 1.4
	m := make(map[int]int, 0)

	for _, a := range A {
		m[a] = 0
	}
	for key, _ := range m {
		for _, a := range A {
			if a == key {
				m[key]++
			}
		}
	}
	var keys []int
	for _, value := range m {
		keys = append(keys, value)
	}
	sort.Ints(keys)

	if isDif(keys) {
		max := keys[len(keys)-1]

		result := 0
		for key, value := range m {
			if value == max {
				result = key
				break
			}
		}
		for i, a := range A {
			if a == result {
				result = i
			}
		}
		return result
	}
	return -1
}