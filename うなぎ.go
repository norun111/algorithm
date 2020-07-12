package GoPaiza

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isFull(startNum int, peopleNum int, slice []int) bool {
	judge := true

	for i := startNum - 1; i < startNum+peopleNum-2; i++ {
		if slice[i] == 1 || i > len(slice) {
			judge = false
		}
	}
	return judge
}

func oneCount(slice []int) int {

	sum := 0
	for _, s := range slice {
		if s == 1 {
			sum++
		}
	}

	return sum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	chairNum, _ := strconv.Atoi(s[0])
	groupNum, _ := strconv.Atoi(s[1])

	slice := make([]int, chairNum)

	for i := 0; i < groupNum; i++ {
		sc.Scan()
		sl := strings.Split(sc.Text(), " ")
		num, _ := strconv.Atoi(sl[0])
		sitStart, _ := strconv.Atoi(sl[1])

		if sitStart+num-2 <= len(slice) {
			if isFull(sitStart, num, slice) {
				for i := sitStart - 1; i <= sitStart+num-2; i++ {
					slice[i] = 1
				}
			}
		}
	}
	fmt.Println(oneCount(slice))

}
