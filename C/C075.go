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
	l := strings.Split(sc.Text(), " ")
	charge, _ := strconv.Atoi(l[0])
	num, _ := strconv.Atoi(l[1])

	point := 0

	for i := 0; i < num; i++ {
		sc.Scan()
		fee, _ := strconv.Atoi(sc.Text())

		if point >= fee {
			point -= fee
			fmt.Println(charge, point)
		} else if charge >= fee {
			charge -= fee
			p := fee / 10
			point += p
			fmt.Println(charge, point)
		}
	}
}
