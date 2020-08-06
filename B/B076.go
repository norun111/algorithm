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
	l := strings.Split(sc.Text(), " ")
	bredNum, _ := strconv.Atoi(l[0])
	orderNum, _ := strconv.Atoi(l[1])

	bredList := []map[string]int{}

	for i:=0; i<bredNum; i++ {
		sc.Scan()
		li := strings.Split(sc.Text(), " ")

		m := make(map[string]int)

		money, _ := strconv.Atoi(li[0])
		amount, _ := strconv.Atoi(li[1])
		m["金額"] = money
		m["個数"] = amount

		bredList = append(bredList, m)
	}

	for i:=0; i<orderNum; i++ {
		sc.Scan()
		orderList := strings.Split(sc.Text(), " ")

		orderType := orderList[0]

		sum := 0

		canBuy := true

		reduceList := []int{}

		if orderType == "buy" {

			for j:=1; j< bredNum + 1; j++ {
				num, _ := strconv.Atoi(orderList[j])

				//在庫がある時
				if bredList[j-1]["個数"] >= num {

					result := num*bredList[j-1]["金額"] // 合計金額
					sum+=result

					reduceList = append(reduceList, num)

				} else {

					canBuy = false

				}
			}
			if canBuy == true {

				fmt.Println(sum)

				for d, r := range reduceList {
					bredList[d]["個数"]-=r
				}

			} else {
				fmt.Println("-1")
			}

		} else if orderType == "bake" {

			for j:=1; j< bredNum + 1; j++ {
				num, _ := strconv.Atoi(orderList[j])

				bredList[j-1]["個数"] += num
			}

		}
	}
}