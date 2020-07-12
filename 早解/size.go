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
	num, _ := strconv.ParseFloat(sc.Text(), 32)

	us := num - 18.0
	uk := num - 18.5
	//小数点第一位まで取得　prec
	fmt.Println(strconv.FormatFloat(us, 'f', 1, 32) + " " + strconv.FormatFloat(uk, 'f', 1, 32))
}
