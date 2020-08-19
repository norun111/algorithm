package main
import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	li := strings.Split(sc.Text(), " ")
	Y, _ := strconv.Atoi(li[0])
	X, _ := strconv.Atoi(li[1])
	D := li[2]

	sc.Scan()
	direction := sc.Text()

	if D == "N" {
		if direction == "R" {
			X++
		} else {
			X--
		}
	} else if D == "S" {
		if direction == "R" {
			X--
		} else {
			X++
		}
	} else if D == "E" {
		if direction == "R" {
			Y++
		} else {
			Y--
		}
	} else if D == "W" {
		if direction == "R" {
			Y--
		} else {
			Y++
		}
	}

	fmt.Println(Y, X)
}