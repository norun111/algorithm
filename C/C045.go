package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// 	"sort"
	"math"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	num, _ := strconv.Atoi(s[0])
	perPage, _ := strconv.Atoi(s[1])
	whichPage, _ := strconv.Atoi(s[2])

	pageNum := float64(num) / float64(perPage)
	// fmt.Printf("%T",pageNum)
	pageNum = math.Ceil(pageNum)
	p := int(pageNum)

	slice := make([][]int, p)

	isWhere := 1
	index := 0

	for i := 1; i <= num; i++ {

		//10*1, 10*2, 10*3....
		if isWhere*perPage == i {
			slice[index] = append(slice[index], i)
			isWhere++
			index++
		} else {
			slice[index] = append(slice[index], i)
		}
	}

	if whichPage > p {
		fmt.Println("none")
	} else {
		result := slice[whichPage-1]

		strSlice := []string{}

		for _, r := range result {
			str := strconv.Itoa(r)
			strSlice = append(strSlice, str)
		}
		fmt.Println(strings.Join(strSlice, " "))
	}
}
