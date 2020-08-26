package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//   	loopNum, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	nums := strings.Split(sc.Text(), " ")
	intNums := make([]int, 0)
	for _, n := range nums {
		i, _ := strconv.Atoi(n)
		intNums = append(intNums, i)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(intNums)))

	alice := 0
	bob := 0

	for i, card := range intNums {
		if i%2==0 {
			alice+=card
		} else {
			bob+=card
		}
	}
	fmt.Println(alice-bob)
}