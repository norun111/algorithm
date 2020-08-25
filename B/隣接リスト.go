package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
	"strings"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	l := strings.Split(sc.Text(), " ")
	node, _ := strconv.Atoi(l[0])
	edge, _ := strconv.Atoi(l[1])

	line := make([][]int, node, node)

	for i:=0; i<edge; i++ {
		sc.Scan()
		l := strings.Split(sc.Text(), " ")

		first, _ := strconv.Atoi(l[0])
		second, _ := strconv.Atoi(l[1])

		line[first-1] = append(line[first-1], second-1)
		line[second-1] = append(line[second-1], first-1)
	}


	for i:=0; i<len(line); i++ {
		sort.Ints(line[i])
	}

	strLine := make([][]string, 0)
	for i:=0; i<len(line); i++ {
		str := make([]string, 0)
		for j:=0; j<len(line[i]); j++ {
			s := strconv.Itoa(line[i][j])
			str = append(str, s)
		}
		strLine= append(strLine, str)
	}

	for _, s := range strLine {
		fmt.Println(strings.Join(s, ""))
	}
}