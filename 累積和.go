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
	num, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	strNums := strings.Split(sc.Text(), " ")
	
	intNums := make([]int, 0, num)
	for _, str := range strNums {
		i, _ := strconv.Atoi(str)
		intNums = append(intNums, i)
	}

	sum := 0
	for _, i := range intNums {
		sum += i
		fmt.Println(sum)
	}
}