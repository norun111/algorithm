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

	result := 0

	for i, _ := range intSlice {
		sum := 0
		for j:=i; j<len(intSlice); j++ {

			sum+=intSlice[j]
			if sum >= M {
				if result==0 || j+1<result {
					result = j+1
				}
				break
			}
		}
	}
	if result == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(result)
	}
}
