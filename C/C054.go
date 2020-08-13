package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	loopNum, _ := strconv.Atoi(s[0])
	speedLimit, _ := strconv.Atoi(s[1])

	obj := make(map[int]int, loopNum)

	for i := 0; i < loopNum; i++ {
		sc.Scan()
		sl := strings.Split(sc.Text(), " ")

		dist, _ := strconv.Atoi(sl[0])
		speed, _ := strconv.Atoi(sl[1])

		obj[dist] = speed
	}

	// 	fmt.Println(obj)

	resultList := []int{}

	for key1 := range obj {
		for key2 := range obj {

			if key1 < key2 {
				dif := key2 - key1
				speedDif := obj[key2] - obj[key1]

				r := speedDif / dif
				resultList = append(resultList, r)
			}
		}
	}
	sin := false

	for _, re := range resultList {
		if speedLimit < re {
			fmt.Println("YES")
			sin = true
			break
		}
	}

	if sin == false {
		fmt.Println("NO")
	}
}
