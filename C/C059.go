package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	slice := make([]int, 4)

	for i:=0; i<loopNum; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), "")
		// fmt.Println(s)
		// fmt.Println("*****************")
		for j:=0; j<4; j++ {
			// fmt.Println(s[j])
			if s[j] == "1" {
				slice[j]+=1
			}
		}
	}

	// 	fmt.Println(slice)

	resultList := []string{}
	for _, s := range slice {
		if s >= 2 {
			resultList = append(resultList, "0")
		} else {
			str := strconv.Itoa(s)
			resultList = append(resultList, str)
		}
	}
	fmt.Println(strings.Join(resultList, ""))
}