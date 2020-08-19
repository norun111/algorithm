package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lastNumber(hoge string) int {
	// 文字列の一文字目をrune型にキャストしています。
	last := rune(hoge[len(hoge)-1])

	// unicodeパッケージのIsUpper関数を使用しています。
	// 引数にrune型の変数initialsをいれ、
	// そのrune文字が大文字かどうかをbool値で返します。
	str := string(last)
	lastInt, _ := strconv.Atoi(str)
	return lastInt
}

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	employeeNum, _ := strconv.Atoi(list[0])
	groupNum, _ := strconv.Atoi(list[1])
	messagesNum, _ := strconv.Atoi(list[2])

	employeeList := make([]int, 0)

	for i:=1; i<=employeeNum; i++ {
		employeeList = append(employeeList, i)
	}

	groupInfoSlice := make([][]int, 0)

	for i:=0; i<groupNum; i++ {
		sc.Scan()
		groupInfo := strings.Split(sc.Text(), " ")
		newGroupInfo := make([]int, 0)

		for _, g := range groupInfo {
			gi, _ := strconv.Atoi(g)
			newGroupInfo = append(newGroupInfo, gi)
		}

		groupInfoSlice = append(groupInfoSlice, newGroupInfo)
	}

	messagesSlice := make([][]string, employeeNum)

	for i:=0; i<messagesNum; i++ {
		sc.Scan()
		messageInfo := strings.Split(sc.Text(), " ")
		from, _ := strconv.Atoi(messageInfo[0])
		group, _ := strconv.Atoi(messageInfo[1])
		to, _ := strconv.Atoi(messageInfo[2])
		message := messageInfo[3]

		if group == 0 {
			messagesSlice[from-1] = append(messagesSlice[from-1], message)
			messagesSlice[to-1] = append(messagesSlice[to-1], message)
		} else if group == 1 {
			groupIndex := lastNumber(message) - 1 // groupInfoSliceのインデックス1なら0番目を指す
			nowGroup := groupInfoSlice[groupIndex]
			groupMemberNum := nowGroup[0]

			for i:=1; i<=groupMemberNum; i++ {
				messagesSlice[nowGroup[i]-1] = append(messagesSlice[nowGroup[i]-1], message)
			}
		}
	}

	for i, message := range messagesSlice {

		for _, m := range message {
			fmt.Println(m)
		}

		if i==len(messagesSlice)-1 {

		} else {
			fmt.Println("--")
		}
	}
}