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
	loopNum, _ := strconv.Atoi(sc.Text())

	//オブジェクトの作成
	obj := make(map[string]map[string]int, 0)

	typeList := []string{}
	RLlist := []string{}

	for i := 0; i < loopNum; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		obj[s[0]] = map[string]int{
			"R": 0,
			"L": 0,
		}
		typeList = append(typeList, s[0])
		RLlist = append(RLlist, s[1])
	}
	// fmt.Println(obj)
	// fmt.Println(typeList)
	// fmt.Println(RLlist)

	for key := range obj {

		for i := 0; i < loopNum; i++ {

			if key == typeList[i] {

				if RLlist[i] == "R" {

					obj[key]["R"] += 1
				} else if RLlist[i] == "L" {
					obj[key]["L"] += 1
				}
			}
		}
	}
	// fmt.Println(obj)

	sum := 0

	for key := range obj {

		add := 0
		add += obj[key]["R"]
		add += obj[key]["L"]

		// fmt.Println(add)
		result := add / 2
		sum += result
	}

	fmt.Println(sum)

}
