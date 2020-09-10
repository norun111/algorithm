package Codity
// 78% runtime
// you can also use imports, for example:
import (
	"math"
	// "fmt"
)
// import "os"
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
	absList := []int{}

	for _, num := range A {
		if !isContain(absList, int(math.Abs(float64(num)))) {
			absList =  append(absList, int(math.Abs(float64(num))))
		}
	}
	return len(absList)
}