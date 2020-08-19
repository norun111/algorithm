package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func next(nowInt int, str string, nowSlice []int) ([]int, int) {

	if nowInt == 6 {
		if str == "U" {
			//上方向
			nowInt = 2
			nowSlice[0]--
		} else if str == "D" {
			//下方向
			// gyouが＋１
			nowInt = 5
			nowSlice[0]++

		} else if str == "L" {
			// 左方向
			nowInt = 3
			nowSlice[1]--
		} else if str == "R" {
			// 右方向
			nowInt = 4
			nowSlice[1]++
		}
	} else if nowInt == 5 {
		if str == "U" {
			nowInt = 6
			nowSlice[0]--
		} else if str == "D" {
			//下方向
			// gyouが＋１
			nowInt = 1
			nowSlice[0]++

		} else if str == "L" {
			// 左方向
			nowInt = 3
			nowSlice[1]--
		} else if str == "R" {
			// 右方向
			nowInt = 4
			nowSlice[1]++
		}
	} else if nowInt == 4 {
		fmt.Println("now")
		if str == "U" {
			nowInt = 6
			nowSlice[0]--
		} else if str == "D" {
			nowInt = 1
			nowSlice[0]++
		} else if str == "L" {
			// 左方向
			nowInt = 5
			nowSlice[1]--
		} else if str == "R" {
			// 右方向

			nowInt = 2
			nowSlice[1]++
		}
	} else if nowInt == 3 {
		if str == "U" {
			nowInt = 2
			nowSlice[0]--
		} else if str == "D" {
			//下方向
			// gyouが＋１
			nowInt = 5
			nowSlice[0]++
		} else if str == "L" {
			// 左方向
			nowInt = 1
			nowSlice[1]--
		} else if str == "R" {
			// 右方向
			nowInt = 6
			nowSlice[1]++
		}
	} else if nowInt == 2 {
		if str == "U" {
			nowInt = 4
			nowSlice[0]--
		} else if str == "D" {
			//下方向
			// gyouが＋１
			nowInt = 3
			nowSlice[0]++
		} else if str == "L" {
			// 左方向
			nowInt = 1
			nowSlice[1]--
		} else if str == "R" {
			// 右方向
			nowInt = 6
			nowSlice[1]++
		}
	} else if nowInt == 1 {
		if str == "U" {
			nowInt = 2
			nowSlice[0]--
		} else if str == "D" {
			//下方向
			// gyouが＋１
			nowInt = 5

		} else if str == "L" {
			// 左方向
			nowInt = 4
			nowSlice[1]--
		} else if str == "R" {
			// 右方向
			nowInt = 3
			nowSlice[1]++
		}
	}

	return nowSlice, nowInt

}


func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	// 	loopNum, _ := strconv.Atoi(list[0])
	tate, _ := strconv.Atoi(list[1])
	yoko, _ := strconv.Atoi(list[2])

	massSlice := make([][]int, 0, tate)

	for i:=0; i<tate; i++ {

		slice := make([]int, yoko, yoko)

		massSlice = append(massSlice, slice)
	}

	sc.Scan()
	li := strings.Split(sc.Text(), " ")
	gyou, _ := strconv.Atoi(li[0])
	letsu, _ := strconv.Atoi(li[1])

	massSlice[gyou-1][letsu-1] = 6

	fmt.Println(massSlice)

	sc.Scan()
	orderList := strings.Split(sc.Text(), "")

	// DRRUULL

	nowInt := 6

	placeList := []int{gyou, letsu}

	for _, order := range orderList {
		fmt.Println(next(nowInt, order, placeList))
	}

}