package main

import (
	"fmt"
	"math"
)

func main() {
	//切り上げ
	i := 0
	v := float64(1.234)
	i = int(math.Ceil(v))
	fmt.Println(i)

	//切り捨て
	i = 0
	v = float64(1.234)
	i = int(math.Floor(v))
	fmt.Println(i)
}
