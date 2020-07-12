package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), " ")
	// fmt.Println(slice)

	m := make(map[string]int, 0)

	for _, s := range slice {
		m[s] = 0
	}
	// fmt.Println(m)

	for key := range m {

		for _, s := range slice {

			if key == s {
				m[key] += 1
			}
		}
	}

	// fmt.Println(m)

	//mapにすると順序が不定なので出力が順番通りになるように新しいスライスを作成
	compareList := make([]string, 0)

	for _, s := range slice {

		for key := range m {

			if s == key {

				// fmt.Println(s,i)
				if !isContain(compareList, s) {
					compareList = append(compareList, s)
				}
			}
		}
	}
	fmt.Println(compareList)

	resultList := make([]int, 0)

	for _, c := range compareList {

		for key := range m {

			if c == key {
				resultList = append(resultList, m[c])
			}
		}
	}
	// fmt.Println(resultList)

	for i, c := range compareList {

		fmt.Println(c, resultList[i])
	}

}

func isContain(slice []string, str string) bool {

	condition := false
	for _, s := range slice {

		if s == str {
			condition = true
			break //一番初めだけ検索する
		}
	}

	return condition
}
