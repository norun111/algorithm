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
	N, _ := strconv.Atoi(li[2])

	for i:=0; i<N; i++ {
		sc.Scan()
		direction := sc.Text()

		if direction == "N" {
			Y--
		} else if direction == "E" {
			X++
		} else if direction == "W" {
			X--
		} else if direction == "S" {
			Y++
		}
		fmt.Println(Y, X)
	}
}