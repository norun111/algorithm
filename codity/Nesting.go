package Codity

// you can also use imports, for example:
import (
	"strings"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func isValidPair(left, right string) bool {
	if left == "(" && right == ")" {
		return true
	}
	return false
}


func Solution(S string) int {
	strSlice := strings.Split(S, "")

	stack := make([]string, 0)

	for _, str := range strSlice {
		if str == "(" {
			stack = append(stack, str)
		} else {
			if len(stack) != 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]// 末尾以外を配列に更新

				if !isValidPair(last, str) {
					return 0
				}
			} else {
				return 0
			}
		}
	}

	if len(stack) == 0 {
		return 1
	}
	return 0
}