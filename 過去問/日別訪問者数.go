package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), " ")
	num, _ := strconv.Atoi(slice[0])
	div, _ := strconv.Atoi(slice[1])

	loopNum := (num - div) + 1

	sc.Scan()
	numSlice := strings.Split(sc.Text(), " ")
	// 	fmt.Println(numSlice)

	var result [][]string

	for i := 0; i < loopNum; i++ {
		//fmt.Println(numSlice[i],numSlice[i+1], numSlice[i+2])

		extra := numSlice[i : div+i]

		result = append(result, extra)
	}

	// 	fmt.Println(result)

	var newResult []int
	for _, r := range result {

		sum := 0
		for _, a := range r {
			i, _ := strconv.Atoi(a)
			sum += i
		}
		newResult = append(newResult, sum)
	}

	t := make([]int, len(newResult))

	copy(t, newResult) //newResultのコピー

	sort.Ints(newResult)

	// 	fmt.Println(newResult)
	//スライスtの中身を文字列に変換する
	li := make([]string, 0)
	for _, a := range t {
		st := strconv.Itoa(a)
		li = append(li, st)
	}

	// 	fmt.Println(li)

	max := newResult[len(newResult)-1]

	maxStr := strconv.Itoa(max)
	in := index(li, maxStr)

	// 	fmt.Println(in)

	resultSum := 0
	for _, r := range newResult {
		if max == r {
			resultSum++
		}
	}

	fmt.Println(resultSum, in+1) //候補日数
}

func index(slice []string, str string) int {

	index := 0
	for i, s := range slice {

		if s == str {
			//  fmt.Println(s, i)
			index = i
			break //
		}
	}

	return index
}
