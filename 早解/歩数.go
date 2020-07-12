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
	slice := strings.Split(sc.Text(), " ")
	km, _ := strconv.Atoi(slice[0])
	km = km * 100000
	hohaba, _ := strconv.Atoi(slice[1])

	num := km / hohaba

	if num >= 10000 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
