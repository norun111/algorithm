//一応できていると思うが問題の意味があまりわからない

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	l := strings.Split(sc.Text(), " ")

	bushesNum, _ := strconv.Atoi(l[0])
	rabbitNum, _ := strconv.Atoi(l[1])
	jumpNum, _ := strconv.Atoi(l[2])

	bushesList := []int{1}

	for i := bushesNum; i > 1; i-- {
		bushesList = append(bushesList, i)
	}

	rabbitPlace := make([]int, bushesNum)

	for i := 0; i < rabbitNum; i++ {
		sc.Scan()
		place, _ := strconv.Atoi(sc.Text())

		for a, b := range bushesList {
			if b == place {
				rabbitPlace[a] = i + 1 //1番目のウサギ 2番目の...
			}
		}
	}

	// fmt.Println(bushesList)

	// fmt.Println(rabbitPlace)

	nowTurn := 1

	count := 0

	for i := 0; i < jumpNum; i++ {
		for index, r := range rabbitPlace {
			if count == jumpNum {
				break
			}
			if r == nowTurn {

				if bushesNum-1 == index {
					rabbitPlace[index] = 0
					rabbitPlace[0] = nowTurn
					count++

				} else {
					rabbitPlace[index] = 0
					rabbitPlace[index+1] = nowTurn
					count++
				}
				if nowTurn+1 > rabbitNum {
					nowTurn = 1
				} else {
					nowTurn++
				}
			}
		}
	}
	// fmt.Println(bushesList)

	// fmt.Println(rabbitPlace)

	m := make([]int, rabbitNum)

	for i, r := range rabbitPlace {

		if r > 0 {
			m[r-1] = bushesList[i]
		}
	}

	// fmt.Println(m)

	for _, result := range m {
		fmt.Println(result)
	}
}
