package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func sum(i int) int {
	str := strconv.Itoa(i)

	sum := 0
	for i:=0; i< len(str); i++ {
		in, _ := strconv.Atoi(string(str[i]))
		sum+=in
	}
	return sum
}

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	numsList := strings.Split(sc.Text(), " ")
	max, _ := strconv.Atoi(numsList[0])
	lower, _ := strconv.Atoi(numsList[1])
	upper, _ := strconv.Atoi(numsList[2])

	nums := make([]int, 0)
	for i:=1; i<=max; i++ {

		if sum(i)>=lower && upper>=sum(i) {
			nums = append(nums, i)
		}
	}

	sum := 0
	for _, n := range nums {
		sum+=n
	}

	fmt.Println(sum)
}