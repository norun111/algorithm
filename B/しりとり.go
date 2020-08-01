package main
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)
func main(){
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	li := strings.Split(sc.Text(), " ")
	peopleNum, _ := strconv.Atoi(li[0])
	wordNum, _ := strconv.Atoi(li[1])
	speakNum, _ := strconv.Atoi(li[2])

	wordList := []string{}

	peopleList := []int{}

	for i:=0; i<wordNum; i++ {
		sc.Scan()
		wordList = append(wordList, sc.Text())
	}

	for i:=1; i<=peopleNum; i++ {
		peopleList = append(peopleList, i)
	}
	// fmt.Println(peopleList)//[1 2 3]
	// fmt.Println(wordList)//[a aloha app az paiza warp]

	nowTurn := 0

	alreadySay := []string{}
	alreadySay2 := []string{}

	for i:=0; i<speakNum; i++ {
		// fmt.Println(peopleList)
		// fmt.Println(alreadySay)
		sc.Scan()
		word := sc.Text()

		wordSplit := strings.Split(word, "")

		if i == 0 {

			if index(wordList, word) == -1 {
				//発言は、単語リストにある K 個の単語のうちのいずれかの単語でなければならない。
				peopleList = unset(peopleList, nowTurn)
			} else if wordSplit[len(wordSplit)-1] == "z" {
				peopleList = unset(peopleList, nowTurn)
			} else {

				//発言した後にリストに追加
				alreadySay = append(alreadySay, word)
				alreadySay2 = append(alreadySay, word)
			}

			nowTurn++ //次の人に進める


		} else {

			//最初の人以外

			beforeWord := alreadySay[len(alreadySay)-1]

			beforeWordSplit := strings.Split(beforeWord, "")

			if len(alreadySay2) == 0 {

				// リストがミスによって初期化された際

				if index(wordList, word) == -1 {
					//発言は、単語リストにある K 個の単語のうちのいずれかの単語でなければならない。
					alreadySay2 = nil
					peopleList = unset(peopleList, nowTurn)

				} else if index(alreadySay, word) != -1 {
					//すでに言っていた場合
					alreadySay2 = nil
					peopleList = unset(peopleList, nowTurn)

				} else if wordSplit[len(wordSplit)-1] == "z" {
					alreadySay2 = nil
					peopleList = unset(peopleList, nowTurn)

				} else {
					//発言した後にリストに追加
					alreadySay = append(alreadySay, word)
					alreadySay2 = append(alreadySay, word)

					if nowTurn + 1 == len(peopleList) {
						nowTurn = 0
					} else {
						nowTurn++
					}
				}

			} else {

				if index(wordList, word) == -1 {
					//発言は、単語リストにある K 個の単語のうちのいずれかの単語でなければならない。
					alreadySay2 = nil
					peopleList = unset(peopleList, nowTurn)

				} else if index(alreadySay, word) != -1 {
					//すでに言っていた場合
					alreadySay2 = nil
					peopleList = unset(peopleList, nowTurn)

				} else if beforeWordSplit[len(beforeWordSplit)-1] != wordSplit[0] {
					alreadySay2 = nil
					peopleList = unset(peopleList, nowTurn)

				} else if wordSplit[len(wordSplit)-1] == "z" {
					alreadySay2 = nil
					peopleList = unset(peopleList, nowTurn)

				} else {
					// fmt.Println(peopleList[nowTurn], word, nowTurn)
					//発言した後にリストに追加
					alreadySay = append(alreadySay, word)
					alreadySay2 = append(alreadySay, word)

					if nowTurn + 1 == len(peopleList) {
						nowTurn = 0
					} else {
						nowTurn++
					}
				}




			}



		}

	}

	fmt.Println(len(peopleList))

	for _, p := range peopleList {
		fmt.Println(p)
	}

}

func index(slice []string, str string) int {

	index := -1
	for i, s := range slice {

		if s == str {
			index = i
			break //一番初めだけ検索する
		}
	}

	return index // indexが-1は配列に文字列が存在しない場合
}

func unset(s []int, i int) []int {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}