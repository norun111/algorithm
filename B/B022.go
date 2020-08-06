//すべてのテストケースは通っていない
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
	list := strings.Split(sc.Text(), " ")
	candidatesNum, _ := strconv.Atoi(list[0])
	votersNum, _ := strconv.Atoi(list[1])
	addressNum, _ := strconv.Atoi(list[2])

	voteList := make([]int, candidatesNum+1)

	voteList[0] = votersNum

	for i := 0; i < addressNum; i++ {
		sc.Scan()
		addresser, _ := strconv.Atoi(sc.Text())
		// fmt.Println(addresser)
		count := 0

		for j := 0; j < len(voteList); j++ {

			if j != addresser && voteList[j] > 0 {
				voteList[j]--
				count++
			}
		}
		voteList[addresser] += count
	}

	voteList = voteList[1:] // 支持者無しは除外

	resultList := []int{voteList[0]}
	indexList := []int{1}

	for i := 1; i < len(voteList); i++ {

		if resultList[len(resultList)-1] <= voteList[i] {
			if resultList[len(resultList)-1] != voteList[i] {
				indexList = indexList[:len(indexList)-1]
			}
			resultList = resultList[:len(resultList)-1]
			resultList = append(resultList, voteList[i])
			indexList = append(indexList, i+1)
		}
	}

	for _, i := range indexList {
		fmt.Println(i)
	}

}
