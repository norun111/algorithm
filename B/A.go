package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isContain(slice []string, str string) []string {

	condition := []string{}

	for _, s := range slice {

		l := strings.Split(s, " ")
		start := l[0]

		if start == str {
			condition = append(condition, s)
		}
	}

	return condition
}

func unset(s []string, i int) []string {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func main(){
	// Your code here!
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	// placeNum, _ := strconv.Atoi(list[0])
	trainNum, _ := strconv.Atoi(list[1])

	trainList := []string{}

	dummyList := []string{}

	for i:=0; i<trainNum; i++ {
		sc.Scan()
		li := sc.Text()
		trainList = append(trainList, li)
		dummyList = append(dummyList, li)
	}

	for i:=0; i<trainNum; i++ {
		fmt.Println(dummyList)

		l := strings.Split(trainList[i], " ")
		// start := l[0]
		end := l[1]

		fmt.Println("************")

		resultList := [][]string{}

		for j:=0; j<len(dummyList); j++ {

			if i != j {
				// fmt.Println(trainList[j])
				anotherTrain := strings.Split(trainList[j], " ")
				// anotherTrainStart := anotherTrain[0]
				// anotherTrainEnd := anotherTrain[1]

				anotherDummyTrain := strings.Split(dummyList[j], " ")
				anotherDummyTrainStart := anotherTrain[0]
				anotherDummyTrainEnd := anotherTrain[1]
				fmt.Println(resultList, anotherDummyTrain)

				if end == anotherDummyTrainStart {
					result := isContain(dummyList, anotherDummyTrainEnd)

					if len(result) > 0 {
						// fmt.Println(isContain(dummyList, anotherTrainEnd))
						// 今調査している要素を削除
						dummyList = unset(dummyList, j)

						resultList = append(resultList, anotherTrain , result)
					}
				}
			}
		}
		fmt.Println(resultList)
	}

}
