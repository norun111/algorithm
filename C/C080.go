package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
func main(){
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	buttonNum,_ := strconv.Atoi(s[0])
	missNum, _ := strconv.Atoi(s[1])

	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	numList := strings.Split(sc.Text(), " ")

	nowNumber := 0
	x := 0
	y := 0

	for i:=0; i<loopNum; i++ {
		number, _ := strconv.Atoi(numList[i])
		if i==0{

			//ミス(1番目だから)
			if number > 1 {
				y++
				nowNumber++
			}

			//最初の番号がボタンの数より多ければ
			//ミス
			if number > buttonNum {
				y++
				nowNumber++
			}

			//あっている時　最初が1
			if number == 1 {
				// fmt.Println(i, number)
				x++
				nowNumber++
			}

		} else {


			//1番目以外の処理

			//順番通り
			if number == nowNumber+1 {
				x++
				nowNumber++
				if nowNumber == buttonNum {
					nowNumber = 0
				}
				// fmt.Println(nowNumber)
				// fmt.Println(i, number)
			} else {
				//間違い
				// fmt.Println("miss")
				// fmt.Println(nowNumber)
				// fmt.Println(i, number)
				y++
				nowNumber++
				if nowNumber == buttonNum {
					nowNumber = 0
				}
				if y == missNum {
					break
				}

			}
		}
	}

	if y == missNum {
		fmt.Println(-1)
	} else {
		fmt.Println(1000*x)
	}
}