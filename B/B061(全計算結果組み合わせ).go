package main

import (
"bufio"
"fmt"
"os"
// 	"sort"
"strconv"
// 	"strings"
"math"
)

func Powerset2(nums []int) [][]int {
	length := int(math.Pow(2, float64(len(nums))))
	result := make([][]int, length)

	index := 0
	result[index] = []int{}
	index++

	for _, n := range nums {
		max := index
		for i := 0; i < max; i++ {
			result[index] = copyAndAppend(result[i], n)
			index++
		}
	}

	return result
}

func copyAndAppend(nums []int, n int) []int {
	dst := make([]int, len(nums)+1)
	copy(dst, nums)
	dst[len(nums)] = n
	return dst
}

func isMin(slice []int, minNum int) bool {

	condition := false

	for i, _ := range slice {
		sum := 0

		for j:=0; j<len(slice); j++ {

			if i != j {
				sum+=slice[j]
			}
		}

		if sum < minNum {
			condition = true
		} else {
			condition = false
			break
		}

	}

	return condition
}


func SumSlice(slice [][]int, minNum int) [][]int {

	newSlice := [][]int{}

	for _, s := range slice {

		sum := 0
		for _, s1 := range s {
			sum+=s1
		}

		if sum >= minNum {
			newSlice = append(newSlice, s)
		}
	}

	return newSlice
}

func main(){
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	minNum, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	slice := make([]int, 0)

	for i:=0; i<loopNum; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())

		slice = append(slice, num)
	}

	t := make([]int, len(slice))

	copy(t, slice)

	resultList := Powerset2(slice)

	newSlice := SumSlice(resultList, minNum)


	sum := 0
	for _, result := range newSlice {
		condition := isMin(result, minNum)

		if condition {
			sum++
		}
	}

	fmt.Println(sum)

}