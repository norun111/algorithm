/*
500 円玉を AA 枚、100 円玉を BB 枚、50 円玉を CC 枚持っています。
これらの硬貨の中から何枚かを選び、合計金額をちょうど
XX 円にする方法は何通りあるでしょうか？
 */

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	coinList := strings.Split(sc.Text(), " ")
	fivehand_coin, _ := strconv.Atoi(coinList[0])
	onehand_coin, _ := strconv.Atoi(coinList[1])
	five_coin, _ := strconv.Atoi(coinList[2])

	count := 0

	for i:=0; i<=fivehand_coin; i++ {
		for j:=0; j<=onehand_coin; j++ {
			for k:=0; k<=five_coin; k++ {
				if 500*i+100*j+50*k == 500 {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}