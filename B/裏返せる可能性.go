package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func playerPlace(massSlice [][]string, X int, Y int) [][]string {

	for i, mass := range massSlice {

		for j, _ := range mass {
			if i == Y {
				massSlice[i][j] = "*"
			}
			if j == X {
				massSlice[i][j] = "*"
			}
		}
	}

	massSlice[Y][X] = "!"

	return massSlice
}

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	li := strings.Split(sc.Text(), " ")
	H, _ := strconv.Atoi(li[0])
	W, _ := strconv.Atoi(li[1])
	Y, _ := strconv.Atoi(li[2])
	X, _ := strconv.Atoi(li[3])

	massSlice := make([][]string, 0, W)

	for i:=0; i<W; i++ {
		slice := make([]string, 0, H)

		for i:=0; i<H; i++ {
			slice = append(slice, ".")
		}

		massSlice = append(massSlice, slice)
	}

	playerPlace(massSlice, X, Y)

	for _, mass := range massSlice {
		fmt.Println(strings.Join(mass, ""))
	}
}