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
	H, _ := strconv.Atoi(li[0])
	W, _ := strconv.Atoi(li[1])

	massSlice := make([][]string, 0)

	for i:=0; i<H; i++ {
		sc.Scan()
		list := strings.Split(sc.Text(), "")
		massSlice = append(massSlice, list)
	}

	for i, mass := range massSlice {

		for j, _ := range mass {
			if W != 1 {
				if j==0 {
					//左端
					if mass[j+1] == "#" {
						fmt.Println(i, j)
					}
				} else if j==W-1 {
					//右端
					if mass[j-1] == "#" {
						fmt.Println(i, j)
					}
				} else {
					if mass[j+1] == "#" && mass[j-1] == "#" {
						fmt.Println(i, j)
					}
				}
			}
		}
	}
}