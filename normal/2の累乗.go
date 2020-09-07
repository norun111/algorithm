package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, _ := strconv.ParseFloat(sc.Text(), 64)
	exp := math.Pow(2, num)
	fmt.Println(exp)
	result := int(exp)%1000000007
	fmt.Println(result)
}