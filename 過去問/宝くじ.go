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
	correct, _ := strconv.Atoi(sc.Text())
	correctStr := sc.Text()

	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	for i := 0; i < loopNum; i++ {
		sc.Scan()
		challenge, _ := strconv.Atoi(sc.Text())
		challengeStr := sc.Text()
		// fmt.Println(charengeStr[len(charengeStr)-3:])
		if challenge == correct {
			fmt.Println("first")
		} else if challenge == correct+1 || challenge == correct-1 {
			fmt.Println("adjacent")
		} else if challengeStr[len(challengeStr)-4:] == correctStr[len(correctStr)-4:] {
			fmt.Println("second")
		} else if challengeStr[len(challengeStr)-3:] == correctStr[len(correctStr)-3:] {
			fmt.Println("third")
		} else {
			fmt.Println("blank")
		}
	}
}
