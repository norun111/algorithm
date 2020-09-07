package Codity
// you can also use imports, for example:
import (
	"strings"
	"fmt"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	bin := fmt.Sprintf("%b", N)

	binSlice := strings.Split(bin, "")
	// fmt.Println(binSlice)

	max_gap_length := 0
	current_gap_length := 0

	flag := false

	for _, b := range binSlice {

		if b == "1" {

			if flag { // 前に1があった時終了する
				flag = false

				if max_gap_length < current_gap_length {
					max_gap_length = current_gap_length
				}
				current_gap_length = 0 // 初期状態に戻す
			} else {
				flag = true
			}
		}
		// flagがtrueの間
		if flag {
			if b == "0" {
				current_gap_length+=1
			}
		}
	}
	// fmt.Println(max_gap_length)
	return max_gap_length
}
