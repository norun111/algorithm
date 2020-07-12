package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	var strikeNum int = 0
	var ballNum int = 0

	for i := 0; i < loopNum; i++ {
		sc.Scan()
		judge := sc.Text()

		if judge == "ball" {
			if ballNum+1 == 4 {
				fmt.Println("fourball!")
				break
			}

			fmt.Println("ball!")
			ballNum++
		} else if judge == "strike" {

			if strikeNum+1 == 3 {
				fmt.Println("out!")
				break
			}
			fmt.Println("strike!")
			strikeNum++
		}
	}
}
