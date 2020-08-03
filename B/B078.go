package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())


	slice := []map[string]int{}

	for i:=0; i<num; i++ {
		m := make(map[string]int)
		sc.Scan()
		list := strings.Split(sc.Text(), " ")
		start, _ := strconv.Atoi(list[0])
		length, _ := strconv.Atoi(list[1])

		m["start"] = start
		m["length"] = length

		slice = append(slice, m)
	}

	var keys []int


	for _, k := range slice {
		keys = append(keys, k["start"])
	}
	sort.Ints(keys)

	slice2 := []map[string]int{}
	for _, k := range keys {
		newMap := make(map[string]int)
		for _, s := range slice {

			if k == s["start"] {
				newMap["start"] = k
				newMap["length"] = s["length"]
			}
		}
		slice2 = append(slice2, newMap)
	}

	endTime := 0
	sum := 0

	for _, s2 := range slice2 {
		st := s2["start"]
		le := s2["length"]

		if st > endTime {
			endTime = st + le
			sum++
		}
	}

	fmt.Println(sum)
}