package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func countOne(list []string) int {
	count := 0

	if len(list) == 0 {
		return 0
	}

	for _, str := range list {
		if str == "1" {
			count++
		}
	}
	return count

}

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), "")

	fmt.Println(countOne(list))
}

// 別解
fmt.Println(strings.Count(a, "1"))