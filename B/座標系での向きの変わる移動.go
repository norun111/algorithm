package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	li := strings.Split(sc.Text(), " ")
	X, _ := strconv.Atoi(li[0])
	Y, _ := strconv.Atoi(li[1])
	N, _ := strconv.Atoi(li[2])

	direction := "N"
	for i:=0; i<N; i++ {
		sc.Scan()
		text := sc.Text()


		if direction == "N" {

			if text == "R" {
				X++
				direction = "E"

			} else {
				X--
				direction = "W"
			}
		} else if direction == "E" {
			if text == "R" {
				Y++
				direction = "S"
			} else {
				Y--
				direction = "N"
			}
		} else if direction == "S" {
			if text == "R" {
				X--
				direction = "W"
			} else {
				X++
				direction = "E"
			}
		} else if direction == "W" {
			if text == "R" {
				Y--
				direction = "N"
			} else {
				Y++
				direction = "S"
			}
		}
		fmt.Println(X, Y)
	}
}