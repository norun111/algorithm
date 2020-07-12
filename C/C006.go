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

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	itemNum, _ := strconv.Atoi(s[0])
	peopleNum, _ := strconv.Atoi(s[1])
	top, _ := strconv.Atoi(s[2])

	sc.Scan()
	sl := strings.Split(sc.Text(), " ")

	pointList := []float64{}
	for _, si := range sl {
		f64, _ := strconv.ParseFloat(si, 64)
		pointList = append(pointList, f64)
	}
	// 	fmt.Println(pointList)

	resultList := []int{}

	for i := 0; i < peopleNum; i++ {
		sc.Scan()
		slice := strings.Split(sc.Text(), " ")

		convertF := []float64{}

		//float64型変換
		for j := 0; j < itemNum; j++ {
			f, _ := strconv.ParseFloat(slice[j], 64)
			convertF = append(convertF, f)
		}

		// fmt.Println(convertF)

		var sum float64 = 0

		for j := 0; j < itemNum; j++ {
			sum += (pointList[j] * convertF[j])
		}

		// fmt.Println(sum)
		v := fmt.Sprintf("%.f", sum)
		result, _ := strconv.Atoi(v)
		resultList = append(resultList, result)
	}
	// fmt.Println(top)
	// fmt.Println(resultList)
	sort.Ints(resultList)

	for i := len(resultList) - 1; i > len(resultList)-(top+1); i-- {
		fmt.Println(resultList[i])
	}
}
