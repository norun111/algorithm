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
	yoko, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	titleList := strings.Split(sc.Text(), " ")

	sc.Scan()
	tate, _ := strconv.Atoi(sc.Text())

	resultList := [][]string{titleList}

	mostLong := make([]int, yoko)

	for i := 0; i < tate; i++ {

		sc.Scan()
		contentList := strings.Split(sc.Text(), " ")

		resultList = append(resultList, contentList)

		for a, c := range contentList {
			if mostLong[a] < len(c) {
				mostLong[a] = len(c)
			} else if mostLong[a] < len(titleList[a]) {
				mostLong[a] = len(titleList[a])
			}
		}
	}
	// 	fmt.Println(mostLong) //各行のLengthの最長リスト
	// 	fmt.Println(resultList)

	resultList2 := [][]string{}

	for i := 0; i < len(resultList); i++ {

		result := []string{}

		for s, r := range resultList[i] {
			//右側につける空白のカウント
			blankNum := mostLong[s] - len(r) + 2 - 1 // 左に一つ空白があるのは確定しているので -1 //+ 2は左右の空白

			// 2行以上だったら
			if yoko > 1 {

				if s == 0 {
					// fmt.Println("|"+" "+ r + strings.Repeat(" ", blankNum) + "|")
					result = append(result, "|"+" "+r+strings.Repeat(" ", blankNum)+"|")
				} else {
					// fmt.Println(" "+ r + strings.Repeat(" ", blankNum) + "|")
					result = append(result, " "+r+strings.Repeat(" ", blankNum)+"|")
				}
			} else {
				// 1行の時
				// fmt.Println("|"+" "+ r + strings.Repeat(" ", blankNum) + "|")
				result = append(result, "|"+" "+r+strings.Repeat(" ", blankNum)+"|")
			}
		}

		resultList2 = append(resultList2, result)
	}

	stickResult := ""

	for i := 0; i < yoko; i++ {

		stickNum := mostLong[i] + 2

		if yoko > 1 {

			if i == 0 {
				// fmt.Println("|" + strings.Repeat("-", stickNum) + "|")
				stickResult += "|" + strings.Repeat("-", stickNum) + "|"
			} else {
				// fmt.Println(strings.Repeat("-", stickNum) + "|")
				stickResult += strings.Repeat("-", stickNum) + "|"
			}

		} else {
			// fmt.Println("|" + strings.Repeat("-", stickNum) + "|")
			stickResult += "|" + strings.Repeat("-", stickNum) + "|"
		}
	}

	fmt.Println(strings.Join(resultList2[0], ""))
	fmt.Println(stickResult)

	for i := 1; i < len(resultList2); i++ {

		fmt.Println(strings.Join(resultList2[i], ""))
	}

}
