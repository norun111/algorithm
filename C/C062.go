package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	// Your code here!
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())

	melon := 0
	sum := 0

	for i:=1; i<num+1; i++ {
		sc.Scan()
		text := sc.Text()

		if text == "melon" {
			sum++
			melon = i
			break
		}
	}


	for i:=melon+1; i<num+1; i++ {
		sc.Scan()
		text := sc.Text()
		if text == "melon" {

			if melon + 10 < i {
				sum++
				melon = i
			}
		}
	}
	fmt.Println(sum)
}
