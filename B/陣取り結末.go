package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func playerPlace(nowPlace []int, massSlice [][]string) []int {
	for i, mass := range massSlice {

		for j, m := range mass {
			if m == "*" {
				nowPlace[0] = i
				nowPlace[1] = j
			}
		}
	}

	return nowPlace
}

func progress(massSlice [][]string, H int, W int) [][]string {
	nowPlace := []int{0, 0}
	nowPlace = playerPlace(nowPlace, massSlice)
	fmt.Println(nowPlace)

	clearSpace := [][]int{}

	//初期値を代入
	clearSpace = append(clearSpace, nowPlace)

	fmt.Println(clearSpace)

	for {
		for i:=0; i<len(clearSpace); i++ {

			y := clearSpace[i][0] //y座標  massSlice[y][x]
			x := clearSpace[i][1] //x座標

			//横幅が1より多い時
			if W > 1 {
				if y == 0 {
					//一番上の時 調べるのは左右と下
					if x == 0 {
						//一番左端
						if massSlice[y][x+1] == "." {
							//右側
							massSlice[y][x+1] = "*"
							s := []int{y, x+1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y+1][x] == "." {
							//下側
							massSlice[y+1][x] = "*"
							s := []int{y+1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}

					} else if x == W-1{
						//一番右
						if massSlice[y][x-1] == "." {
							//左側
							massSlice[y][x-1] = "*"
							s := []int{y, x-1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y+1][x] == "." {
							//下側
							massSlice[y+1][x] = "*"
							s := []int{y+1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
					} else {
						//真ん中
						if massSlice[y][x-1] == "." {
							//左側
							massSlice[y][x-1] = "*"
							s := []int{y, x-1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x+1] == "." {
							//右側
							massSlice[y][x+1] = "*"
							s := []int{y, x+1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y+1][x] == "." {
							//下側
							massSlice[y+1][x] = "*"
							s := []int{y+1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
					}
				} else if y == H-1 {
					//一番下の時 調べるのは上と左右
					if x == 0 {
						// 一番左の時
						if massSlice[y-1][x] == "." {
							//上側
							massSlice[y-1][x] = "*"
							s := []int{y-1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x+1] == "." {
							//右側
							massSlice[y][x+1] = "*"
							s := []int{y, x+1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}

					} else if x == W-1 {
						// 一番右の時
						if massSlice[y-1][x] == "." {
							//上側
							massSlice[y-1][x] = "*"
							s := []int{y-1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x-1] == "." {
							//左側
							massSlice[y][x-1] = "*"
							s := []int{y, x-1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
					} else {
						// 真ん中の時
						if massSlice[y-1][x] == "." {
							//上側
							massSlice[y-1][x] = "*"
							s := []int{y-1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x-1] == "." {
							//左側
							massSlice[y][x-1] = "*"
							s := []int{y, x-1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x+1] == "." {
							//右側
							massSlice[y][x+1] = "*"
							s := []int{y, x+1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
					}
				} else {
					// 真ん中の時なので上下左右調べる
					if x==0 {
						//一番右の時調べるのは上下右
						if massSlice[y-1][x] == "." {
							//上側
							massSlice[y-1][x] = "*"
							s := []int{y-1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y+1][x] == "." {
							//下側
							massSlice[y+1][x] = "*"
							s := []int{y+1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x+1] == "." {
							//右側
							massSlice[y][x+1] = "*"
							s := []int{y, x+1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
					} else if x == W-1 {
						if massSlice[y-1][x] == "." {
							//上側
							massSlice[y-1][x] = "*"
							s := []int{y-1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y+1][x] == "." {
							//下側
							massSlice[y+1][x] = "*"
							s := []int{y+1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x-1] == "." {
							//左側
							massSlice[y][x-1] = "*"
							s := []int{y, x-1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
					} else {
						if massSlice[y-1][x] == "." {
							//上側
							massSlice[y-1][x] = "*"
							s := []int{y-1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y+1][x] == "." {
							//下側
							massSlice[y+1][x] = "*"
							s := []int{y+1, x}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x-1] == "." {
							//左側
							massSlice[y][x-1] = "*"
							s := []int{y, x-1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
						if massSlice[y][x+1] == "." {
							//右側
							massSlice[y][x+1] = "*"
							s := []int{y, x+1}
							clearSpace = append(clearSpace, s) // 進んだ場所を記録
						}
					}
				}
			}

		}
	}

	return massSlice
}

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

	fmt.Println(progress(massSlice, H, W))
}