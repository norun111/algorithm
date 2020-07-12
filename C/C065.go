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
	loopNum, _ := strconv.Atoi(sc.Text())

	maxList := make([]int, 0)
	minList := make([]int, 0)

	divList := make([]int, 0)

	//最大と最小がいくらか
	for i := 0; i < loopNum; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")

		str := s[0]
		num, _ := strconv.Atoi(s[1])
		//以上のもの
		if str == ">" {
			minList = append(minList, num)
		} else if str == "<" {
			maxList = append(maxList, num)
		} else if str == "/" {
			divList = append(divList, num)
		}
	}

	sort.Ints(maxList)
	sort.Ints(minList)

	max := maxList[len(maxList)-1] //最大値
	min := minList[0]              //最小値

	// 	fmt.Println(divList)

	guessList := make([]int, 0)

	for i := min + 1; i < max; i++ {
		guessList = append(guessList, i)
	}

	// 	fmt.Println(guessList)

	result := 0

	for i := 0; i < len(guessList); i++ {

		isAdd := 0

		for j := 0; j < len(divList); j++ {

			if guessList[i]%divList[j] != 0 {
				// fmt.Println(divList[j], j)
			} else {
				// fmt.Println(guessList[i])
				isAdd++
			}
		}

		//"/"が2個の時2個ともで割り切れないといけないので isAddedの数とlen(divList)の数
		if isAdd == len(divList) {
			result = guessList[i]
		}
	}

	fmt.Println(result)
}
