package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	team := list[0]
	passerIndex, _ := strconv.Atoi(list[1])

	sc.Scan()
	teamAList := strings.Split(sc.Text(), " ")
	sc.Scan()
	teamBList := strings.Split(sc.Text(), " ")

	if team == "A" {
		passer, _ := strconv.Atoi(teamAList[passerIndex-1])

		teamB := []int{}
		for _, b := range teamBList {
			intB, _ := strconv.Atoi(b)
			teamB = append(teamB, intB)
		}

		teamA := []int{}
		for _, a := range teamAList {
			intA, _ := strconv.Atoi(a)
			teamA = append(teamA, intA)
		}

		sort.Ints(teamB)
		secondBack := teamB[len(teamB)-2]

		isOffSide := false

		if passer < secondBack {

			for i, ta := range teamA {
				if secondBack < ta {
					fmt.Println(i + 1)
					isOffSide = true
				}
			}
		}
		if isOffSide == false {
			fmt.Println("None")
		}

	} else if team == "B" {
		passer, _ := strconv.Atoi(teamBList[passerIndex-1])

		teamB := []int{}
		for _, b := range teamBList {
			intB, _ := strconv.Atoi(b)
			teamB = append(teamB, intB)
		}

		teamA := []int{}
		for _, a := range teamAList {
			intA, _ := strconv.Atoi(a)
			teamA = append(teamA, intA)
		}

		sort.Ints(teamA)
		// fmt.Println(teamA)
		// fmt.Println(teamB)

		secondBack := teamA[1]
		// fmt.Println(secondBack)

		isOffSide := false

		if passer > secondBack {

			for i, tb := range teamB {
				if secondBack > tb {
					fmt.Println(i + 1)
					isOffSide = true
				}
			}
		}
		if isOffSide == false {
			fmt.Println("None")
		}
	}
}
