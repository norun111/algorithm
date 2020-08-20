package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PlayerPlace(massSlice [][]string, H int, W int) [][]string {

	count := 0

	for i, mass := range massSlice {

		for j, m := range mass {

			if count == 1 {
				break
			}

			if m == "*" {
				if i == 0 {
					//一番上かつ
					if j == 0 {
						//左端
						massSlice[i][j+1] = "*" //右隣を"*"
						massSlice[i+1][j] = "*" //真下を"*"
						count++
					} else if j == W-1{
						//右端
						massSlice[i][j-1] = "*" //左隣を"*"
						massSlice[i+1][j] = "*" //真下を"*"
						count++
					} else {
						//真ん中
						massSlice[i][j+1] = "*" //右隣を"*"
						massSlice[i][j-1] = "*" //左隣を"*"
						massSlice[i+1][j] = "*" //真下を"*"
						count++
					}
				} else if i == H-1 {
					//一番下かつ
					if j == 0 {
						//左端
						massSlice[i-1][j] = "*" //真上を"*"
						massSlice[i][j+1] = "*" //右隣を"*"
						count++
					} else if j == W-1{
						//右端
						massSlice[i-1][j] = "*" //真上を"*"
						massSlice[i][j-1] = "*" //左隣を"*"
						count++
					} else {
						//真ん中
						massSlice[i][j+1] = "*" //右隣を"*"
						massSlice[i][j-1] = "*" //左隣を"*"
						massSlice[i-1][j] = "*" //真上を"*"
						count++
					}
				} else {
					if j == 0 {
						//左端
						massSlice[i+1][j] = "*" //真下を"*"
						massSlice[i-1][j] = "*" //真上を"*"
						massSlice[i][j+1] = "*" //右隣を"*"
						count++
					} else if j == W-1{
						//右端
						massSlice[i+1][j] = "*" //真下を"*"
						massSlice[i-1][j] = "*" //真上を"*"
						massSlice[i][j-1] = "*" //左隣を"*"
						count++
					} else {
						//真ん中
						massSlice[i][j+1] = "*" //右隣を"*"
						massSlice[i][j-1] = "*" //左隣を"*"
						massSlice[i-1][j] = "*" //真上を"*"
						massSlice[i+1][j] = "*" //真下を"*"
						count++
					}

				}
			}

		}
	}

	return massSlice
}
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	li := strings.Split(sc.Text(), " ")
	H, _ := strconv.Atoi(li[0])
	W, _ := strconv.Atoi(li[1])

	massSlice := make([][]string, 0)

	for i:=0; i<H; i++ {
		sc.Scan()
		list := strings.Split(sc.Text(), "")

		massSlice = append(massSlice, list)
	}

	massSlice = PlayerPlace(massSlice, H, W)

	for _, mass := range massSlice {
		fmt.Println(strings.Join(mass, ""))
	}
}