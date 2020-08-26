/*
このままでは調査した部分が全て"*"になるので
もし"*"に行き着いたらそこの座標を調べて
そこまでだけ"*"に塗り替えるという関数を作る必要がある。
 */


package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func playerPlace(massSlice [][]string, H, W, y, x int ) [][]string {
	if y == 0 {
		//一番上の行
		if x == 0 {
			// かつ左端

			// 横の行を調査
			for i, m := range massSlice[0] {
				if i > 0 {
					if m == "*" {
						break
					} else {
						massSlice[0][i] = "*"
					}
				}
			}
			// 縦の行を調査
			for i:=0; i<W; i++ {
				if i > 0 {
					if massSlice[i][0] == "*" {
						break
					} else {
						massSlice[i][0] = "*"
					}
				}
			}
		} else if x == W-1{
			// かつ右端

			// 横の行を調査
			for i:=x; i>=0; i-- {
				if i < x {
					if massSlice[0][i] == "*" {
						break
					} else {
						massSlice[0][i] = "*"
					}
				}
			}
			// 縦の行を調査
			for i:=0; i<W; i++ {
				if i > 0 {
					if massSlice[i][0] == "*" {
						break
					} else {
						massSlice[i][0] = "*"
					}
				}
			}

		} else {
			// かつ真ん中
			// 右横の行を調査
			for i:=x; i<H; i++ {
				if i > x {
					if massSlice[0][i] == "*" {
						break
					} else {
						massSlice[0][i] = "*"
					}
				}
			}
			// 左横の行を調査
			for i:=x; i>=0; i-- {
				if i < x {
					if massSlice[0][i] == "*" {
						break
					} else {
						massSlice[0][i] = "*"
					}
				}
			}
			// 縦の行を調査
			for i:=0; i<W; i++ {
				if i > 0 {
					if massSlice[i][0] == "*" {
						break
					} else {
						massSlice[i][0] = "*"
					}
				}
			}

		}
	} else if y == H-1 {
		// 一番下の行
		if x == 0 {
			// かつ左端
			// 右横の行を調査
			for i, m := range massSlice[y] {
				if i > 0 {
					if m == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}
			// 縦上の列を調査
			for i:=y; i>=0; i-- {
				if i < y {
					if massSlice[i][0] == "*" {
						break
					} else {
						massSlice[i][0] = "*"
					}
				}
			}

		} else if x == W-1{
			// かつ右端
			// 左横の行を調査
			for i:=x; i>=0; i-- {
				if i < x {
					if massSlice[y][i] == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}

			// 縦上の列を調査
			for i:=y; i>=0; i-- {
				if i < y {
					if massSlice[i][x] == "*" {
						break
					} else {
						massSlice[i][x] = "*"
					}
				}
			}
		} else {
			// かつ真ん中

			// 右横の行を調査
			for i:=x; i<H; i++ {
				if i > x {
					if massSlice[y][i] == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}
			// 左横の行を調査
			for i:=x; i>=0; i-- {
				if i < x {
					if massSlice[y][i] == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}
			// 縦上の列を調査
			for i:=y; i>=0; i-- {
				if i < y {
					if massSlice[i][x] == "*" {
						break
					} else {
						massSlice[i][x] = "*"
					}
				}
			}
		}
	} else {
		// 真ん中の行
		if x == 0 {
			// かつ左端 上、 下、 右を調査
			// 縦上の列を調査
			for i:=y; i>=0; i-- {

				if i < y {
					if massSlice[i][x] == "*" {
						break
					} else {
						massSlice[i][x] = "*"
					}
				}
			}
			// 縦下の行を調査
			for i:=0; i<W; i++ {
				if i > y {
					if massSlice[i][0] == "*" {
						break
					} else {
						massSlice[i][0] = "*"
					}
				}
			}
			// 右横の行を調査
			for i, m := range massSlice[y] {
				if i > 0 {
					if m == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}

		} else if x == W-1{
			// かつ右端
			// 左横の行を調査
			for i:=x; i>=0; i-- {
				if i < x {
					if massSlice[y][i] == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}
			// 縦上の列を調査
			for i:=y; i>=0; i-- {
				if i < y {
					if massSlice[i][x] == "*" {
						break
					} else {
						massSlice[i][x] = "*"
					}
				}
			}
			// 縦下の行を調査
			for i:=y; i<W; i++ {
				if i > y {
					if massSlice[i][x] == "*" {
						break
					} else {
						massSlice[i][x] = "*"
					}
				}
			}
		} else {
			// かつ真ん中

			// 右横の行を調査
			for i:=x; i<H; i++ {
				fmt.Println("right")
				if i > x {
					if massSlice[y][i] == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}
			// 左横の行を調査
			for i:=x; i>=0; i-- {
				fmt.Println("left")
				if i < x {
					if massSlice[y][i] == "*" {
						break
					} else {
						massSlice[y][i] = "*"
					}
				}
			}
			// 縦上の列を調査
			for i:=y; i>=0; i-- {
				fmt.Println("up")
				if i < y {
					if massSlice[i][x] == "*" {
						break
					} else {
						massSlice[i][x] = "*"
					}
				}
			}
			// 縦下の行を調査
			for i:=y; i<W; i++ {

				if i > y {
					if massSlice[i][x] == "*" {
						place := []int{i, x}

						break
					} else {

						massSlice[i][x] = "*"
					}
				}
			}
		}
	}

	massSlice[y][x] = "*"
	return massSlice
}
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	H, _ := strconv.Atoi(list[0])
	W, _ := strconv.Atoi(list[1])
	y, _ := strconv.Atoi(list[2])
	x, _ := strconv.Atoi(list[3])

	massSlice := [][]string{}

	for i:=0; i<H; i++ {
		sc.Scan()
		mass := strings.Split(sc.Text(), "")
		massSlice = append(massSlice, mass)
	}

	playerPlace(massSlice, H, W, y, x)

	for _, mass := range massSlice {
		fmt.Println(strings.Join(mass, ""))
	}
}
