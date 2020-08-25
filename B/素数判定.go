package main
import (
	"fmt"
	"os"
	"math"
	"bufio"
	"strconv"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	num, _ := strconv.Atoi(sc.Text())
	var maxNum float64 = float64(num)

	isPrime := true

	for i:=2; i<=int(math.Sqrt(maxNum)) ;i++{
		if int(maxNum)%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime && num != 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}