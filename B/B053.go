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
	list := strings.Split(sc.Text(), " ")
	vertical, _ := strconv.Atoi(list[0])
	horizon, _ := strconv.Atoi(list[1])
	sc.Scan()
	flist := strings.Split(sc.Text(), " ")

	a, _ := strconv.Atoi(flist[0])
	b, _ := strconv.Atoi(flist[1])

	b_a := b - a

	sc.Scan()
	slist := strings.Split(sc.Text(), " ")
	c, _ := strconv.Atoi(slist[0])
	d, _ := strconv.Atoi(slist[1])

	d_c := d - c

	// fmt.Println(b_a, d_c)

	firstList := []int{a, b}
	secondList := []int{c, d}


	for i:=0; i<horizon-2; i++ {
		add1 := firstList[len(firstList)-1] + b_a
		add2 := secondList[len(secondList)-1] + d_c

		firstList = append(firstList, add1)
		secondList = append(secondList, add2)
	}

	saList := []int{}

	for i:=0; i<len(firstList); i++ {
		sa := secondList[i] - firstList[i]
		saList = append(saList, sa)
	}

	// fmt.Println(firstList, secondList, saList)

	resultList := [][]int{}

	resultList = append(resultList, secondList)

	for i:=0; i<vertical-2; i++ {
		l := []int{}

		compareList := resultList[len(resultList)-1]

		for j:=0; j<len(saList); j++ {
			r := compareList[j] + saList[j]
			l = append(l, r)
		}

		resultList = append(resultList, l)
	}

	// fmt.Println(resultList)

	strResultList := [][]string{}

	for _, re := range resultList {
		strl := []string{}

		for j:=0; j<len(saList); j++ {

			s := strconv.Itoa(re[j])
			strl = append(strl, s)
		}
		strResultList = append(strResultList, strl)
	}

	strFirstList := []string{}
	for _, fir := range firstList {
		fi := strconv.Itoa(fir)
		strFirstList = append(strFirstList, fi)
	}

	fmt.Println(strings.Join(strFirstList, " "))

	for _, st := range strResultList {
		fmt.Println(strings.Join(st, " "))
	}
}