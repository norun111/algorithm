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
	nums := strings.Split(sc.Text(), " ")
	start, _ := strconv.Atoi(nums[0])
	end, _ := strconv.Atoi(nums[1])

	sum := 0
	for i:=start; i<=end; i++ {
		sum+=i
	}
	fmt.Println(sum)
}