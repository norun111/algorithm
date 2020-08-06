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
	list := strings.Split(sc.Text(), " ")
	// 	treeNum, _ := strconv.Atoi(list[0])
	averageNum, _ := strconv.Atoi(list[1])

	sc.Scan()
	strlightNumList := strings.Split(sc.Text(), " ")

	lightNumList := []int{}
	for _, s := range strlightNumList {
		intS, _ := strconv.Atoi(s)
		lightNumList = append(lightNumList, intS)
	}

	sc.Scan()
	surveyNum, _ := strconv.Atoi(sc.Text())

	for i := 0; i < surveyNum; i++ {

		sc.Scan()
		distanceList := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(distanceList[0])
		b, _ := strconv.Atoi(distanceList[1])

		sum := 0
		count := 0

		for j := a - 1; j < b; j++ {
			sum += lightNumList[j]
			count++
		}
		result := sum / count

		//平均以下であれば
		if result <= averageNum {
			sub := averageNum - result

			for j := a - 1; j < b; j++ {
				lightNumList[j] += sub
			}
		}
	}

	resultList := []string{}

	for _, l := range lightNumList {
		re := strconv.Itoa(l)
		resultList = append(resultList, re)
	}

	fmt.Println(strings.Join(resultList, " "))
}
