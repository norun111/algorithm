package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func unset(s []string, i int) []string {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func strPlace(slice []string, str string) int {

	index := 0
	for i, s := range slice {

		if s == str {
			index = i
		}
	}

	return index + 1
}

func main() {
	// Your code here!
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	slice := make([]string, 0)

	for i := 0; i < loopNum; i++ {
		sc.Scan()

		str := sc.Text()
		index := strPlace(slice, str)

		//入っている場合
		if arrayContains(slice, str) {
			slice = append([]string{str}, slice...)
			// fmt.Println(slice)
			// fmt.Println(index, str)
			slice = unset(slice, index)

		} else {
			slice = append([]string{str}, slice...)
		}
	}

	// fmt.Println(slice)

	for _, s := range slice {
		fmt.Println(s)
	}
}
