package Codity

// 100%

import (
	"strings"
	// "fmt"
)

func isValidPair(left, right string) bool {
	if left == "(" && right == ")" {
		return true
	} else if left == "[" && right == "]"{
		return true
	} else if left == "{" && right == "}"{
		return true
	}
	return false
}

func Solution(S string) int {
	strSlice := strings.Split(S, "")

	stack := make([]string, 0)
	for _, str := range strSlice {
		if str == "[" || str == "{" || str == "(" {
			stack = append(stack, str)
		} else {
			if len(stack) == 0 {
				return 0
			}
			// stackに"["か"{"か"("が存在する場合
			last := stack[len(stack)-1]  // 末尾を取得
			stack = stack[:len(stack)-1] // 末尾以外を更新

			if !isValidPair(last, str) {
				return 0
			}

		}
	}
	if len(stack) != 0 {
		return 0
	} else {
		return 1
	}
}