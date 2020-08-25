package main
import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"strconv"
	"strings"
)
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	first, _ := strconv.Atoi(list[0])
	second, _ := strconv.Atoi(list[1])

	intList := []int{first, second}
	sort.Ints(intList)

	first = intList[0] // 小さい方
	second = intList[1] // 大きい方

	margin := 0

	for {
		if second%first == 0 {
			break
		}
		margin = second%first
		second = first
		first = margin
	}
	fmt.Println(first)
}