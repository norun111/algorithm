//　まだ完成していない

package main
import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"strconv"
	"strings"
)

type Number struct {
	Count int
	Sum int
}

type Numbers []Number

func (n Numbers) Len() int {
	return len(n)
}

func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Numbers) Less(i, j int) bool {
	if n[i].Count > n[j].Count {
		return n[i].Count > n[j].Count
	}
	return false
}

func (n Numbers) reduce(i int, M int) {

}

func NewMedal(v []int) Medal {
	return Medal{v[0], v[1], v[2]}
}

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), " ")
	// N, _ := strconv.Atoi(slice[0])
	M, _ := strconv.Atoi(slice[1])

	sc.Scan()
	strSlice := strings.Split(sc.Text(), " ")

	intSlice := []int{}
	for _, s := range strSlice {
		i, _ := strconv.Atoi(s)
		intSlice = append(intSlice, i)
	}
	result := []Number{}


	for i, _ := range intSlice {
		sum := 0
		count := 0
		for j:=i; j<len(intSlice); j++ {
			count++
			sum+=intSlice[j]
			if j != len(intSlice)-1 {
				if sum + intSlice[j+1] >= M {
					fmt.Println(sum)
					break
				} else if sum < M {
					fmt.Println(sum)
				}
			} else {
				if intSlice[j] >= M {
					result = append(result, 0)
				}
			}
		}
		fmt.Println("************", count, sum)
		result = append(result, count)
	}

	fmt.Println(result)
	sort.Ints(result)
}