//全て通っていない

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
	loopNum, _ := strconv.Atoi(sc.Text())

	tateList := make([][]int, loopNum)

	numCorrect := []int{}

	for i := 0; i < loopNum; i++ {
		sum := 0
		sc.Scan()
		s := strings.Split(sc.Text(), " ")

		for j := 0; j < loopNum; j++ {

			n, _ := strconv.Atoi(s[j])
			tateList[j] = append(tateList[j], n)
			sum += n
		}
		numCorrect = append(numCorrect, sum)
	}
	sort.Ints(numCorrect)
	correct := numCorrect[len(numCorrect)-1]

	for _, t := range tateList {

		tateSum := 0
		zeroIndex := 0

		for i := 0; i < loopNum; i++ {

			tateSum += t[i]
			if t[i] == 0 {

				zeroIndex = i
			}
		}

		if tateSum < correct {
			diff := correct - tateSum
			t[zeroIndex] = diff
		}
	}
	resultList := make([][]string, loopNum)

	// 	fmt.Println(tateList)

	for _, t := range tateList {

		for j := 0; j < loopNum; j++ {
			str := strconv.Itoa(t[j])
			resultList[j] = append(resultList[j], str)
		}
	}
	// 	fmt.Println(resultList)

	for _, r := range resultList {
		fmt.Println(strings.Join(r, " "))
	}

}
