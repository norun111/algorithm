package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")

	taxiNum, _ := strconv.Atoi(list[0])
	distance, _ := strconv.Atoi(list[1])

	feeList := []int{}

	for i := 0; i < taxiNum; i++ {
		sc.Scan()
		li := strings.Split(sc.Text(), " ")
		firstDistance, _ := strconv.Atoi(li[0])
		firstFee, _ := strconv.Atoi(li[1])
		addDistance, _ := strconv.Atoi(li[2])
		addFee, _ := strconv.Atoi(li[3])

		sum := 0

		if distance >= firstDistance {
			sum += firstFee
			remainingDistance := distance - firstDistance

			if remainingDistance <= addDistance {
				sum += addFee
			} else if remainingDistance%addDistance == 0 {
				div := (float64(remainingDistance) / float64(addDistance))
				in := int(math.Ceil(div))
				sum += (in * addFee)
				sum += addFee
			} else {
				div := (float64(remainingDistance) / float64(addDistance))
				in := int(math.Ceil(div))
				sum += (in * addFee)
			}

		} else if distance < firstDistance {
			sum += firstFee
		}
		feeList = append(feeList, sum)
	}

	sort.Ints(feeList)

	min := strconv.Itoa(feeList[0])
	max := strconv.Itoa(feeList[len(feeList)-1])

	l := []string{min, max}

	fmt.Println(strings.Join(l, " "))

}
