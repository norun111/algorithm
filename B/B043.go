//37点
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	tate, _ := strconv.Atoi(list[0])
	yoko, _ := strconv.Atoi(list[1])

	sc.Scan()
	list2 := strings.Split(sc.Text(), " ")
	x, _ := strconv.Atoi(list2[0])
	y, _ := strconv.Atoi(list2[1])

	coordinate := []int{x, y}

	townList := [][]string{}

	for i := 0; i < tate; i++ {
		sc.Scan()
		list3 := strings.Split(sc.Text(), "")

		town := []string{}

		for _, l := range list3 {
			town = append(town, l)
		}

		townList = append(townList, town)
	}

	count := 0

	whichDir := "right"

	for {
		count++

		if count == 1 {
			//直進するだけ
			if townList[coordinate[0]-1][coordinate[1]-1] == "*" {

				townList[coordinate[0]-1][coordinate[1]-1] = "."
				if coordinate[1]+1 > yoko {
					break
				}
				coordinate[1] += 1

			} else if townList[coordinate[0]-1][coordinate[1]-1] == "." {
				// 庶民の家だった場合
				townList[coordinate[0]-1][coordinate[1]-1] = "*"
				if coordinate[1]+1 > yoko {
					break
				}
				coordinate[1]++
			}
		} else {
			if whichDir == "right" {
				if townList[coordinate[0]-1][coordinate[1]-1] == "*" {
					//富豪であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "."
					coordinate[0]--
					whichDir = "up"
				} else {
					//庶民であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "*"
					if coordinate[0]+1 > tate {
						break
					}
					coordinate[0]++
					whichDir = "down"
				}
			} else if whichDir == "down" {
				if townList[coordinate[0]-1][coordinate[1]-1] == "*" {
					//富豪であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "."
					if coordinate[1]+1 > yoko {
						break
					}
					coordinate[1]++
					whichDir = "right"
				} else {
					//庶民であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "*"
					coordinate[1]--
					whichDir = "left"
				}
			} else if whichDir == "left" {
				if townList[coordinate[0]-1][coordinate[1]-1] == "*" {
					//富豪であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "."
					if coordinate[0]+1 > tate {
						break
					}
					coordinate[0]++
					whichDir = "down"
				} else {
					//庶民であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "*"
					coordinate[0]--
					whichDir = "up"
				}
			} else if whichDir == "up" {
				if townList[coordinate[0]-1][coordinate[1]-1] == "*" {
					//富豪であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "."
					coordinate[1]--
					whichDir = "left"
				} else {
					//庶民であった場合
					townList[coordinate[0]-1][coordinate[1]-1] = "*"
					if coordinate[1]+1 > yoko {
						break
					}
					coordinate[1]++
					whichDir = "right"
				}
			}
		}
	}

	for _, t := range townList {
		fmt.Println(strings.Join(t, ""))
	}
}
