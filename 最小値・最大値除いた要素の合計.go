package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	stringSlice := strings.Split(sc.Text(), " ")

	intSlice := make([]int, 0)

	for _, str := range stringSlice {
		i, _ := strconv.Atoi(str)
		intSlice = append(intSlice, i)
	}

	sort.Ints(intSlice)

	intSlice = intSlice[:len(intSlice)-1]
	intSlice = intSlice[1:]

	sum := 0

	for _, i := range intSlice {
		sum+=i
	}

	//最小値と最大値を除いた要素の合計値
	fmt.Println(sum)
}
