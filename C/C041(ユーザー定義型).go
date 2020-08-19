package main
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Medal struct {
	Gold int
	Silver int
	Bronze int
}

type Medals []Medal

func (m Medals) Len() int {
	return len(m)
}

func (m Medals) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Medals) Less(i, j int) bool {

	if m[i].Gold > m[j].Gold {
		return m[i].Gold > m[j].Gold

	} else if m[i].Gold == m[j].Gold {

		if m[i].Silver > m[j].Silver {
			return m[i].Silver > m[j].Silver
		} else if m[i].Silver == m[j].Silver {

			if m[i].Bronze > m[j].Bronze {
				return m[i].Bronze > m[j].Bronze
			}
		}
	}
	return false
}

//このやり方あり
type Medal struct{ Gold, Silver, Bronze int }

func NewMedal(v []int) Medal {
	return Medal{v[0], v[1], v[2]}
}
//

func main() {
	s := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	medals := []Medal{}
	for _, v := range s {
		medals = append(medals, NewMedal(v))
	}
	fmt.Printf("%#v\n", medals)
}


func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	loopNum, _ := strconv.Atoi(sc.Text())

	var medals Medals = make([]Medal, 0)

	for i:=0; i<loopNum; i++ {
		sc.Scan()
		medalList := strings.Split(sc.Text(), " ")

		//構造体を初期化
		medal := Medal{}
		for i, m := range medalList {
			m, _ := strconv.Atoi(m)
			//各要素に代入
			if i==0 {
				medal.Gold = m
			} else if i==1 {
				medal.Silver = m
			} else if i==2 {
				medal.Bronze = m
			}
		}
		medals = append(medals, medal)
	}
	sort.Sort(medals)

	for _, medal := range medals {
		fmt.Println(medal.Gold, medal.Silver, medal.Bronze)
	}
}