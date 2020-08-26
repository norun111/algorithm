package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	mochiList := make([]int, 0)

	for i:=0; i<loopNum; i++ {
		sc.Scan()
		diameter, _ := strconv.Atoi(sc.Text())
		mochiList = append(mochiList, diameter)
	}
	sort.Ints(mochiList)

	count := 0
	nowMochi := 0
	for _, m := range mochiList {
		if nowMochi < m {
			count++
			nowMochi = m
		}
	}

	fmt.Println(count)
}