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

	massSlice := make([][]string, 0)

	for i:=0; i<H; i++ {
		sc.Scan()
		list := strings.Split(sc.Text(), "")
		massSlice = append(massSlice, list)
	}

	for i, mass := range massSlice {

		for j, _ := range mass {
			if H != 1 {
				if i==0 {
					//上端
					if massSlice[i+1][j] == "#" {
						fmt.Println(i, j)
					}
				} else if i==H-1 {
					//下端
					if massSlice[i-1][j] == "#" {
						fmt.Println(i, j)
					}
				} else {
					if massSlice[i-1][j] == "#" && massSlice[i+1][j]  == "#" {
						fmt.Println(i, j)
					}
				}
			}
		}
	}
}