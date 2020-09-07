package Codity
// Timeout Error 30%

// you can also use imports, for example:
import (
	"math"
	// "fmt"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	max := math.MaxFloat64
	// fmt.Println(max)

	minIndex := 0

	for i:=0; i<len(A); i++ {

		for j:=i+1; j<len(A); j++ {
			// fmt.Println(A[i], A[j])
			//A[i]とA[j]二つの合計の平均値を求める
			var avg1 float64 = (float64(A[i])+float64(A[j]))/2

			if avg1 < max {
				max = avg1
				minIndex = i
			}

			intSlice := []int{}
			for k:=i; k<=j; k++ {
				// 開始位置は上と同じで上は2個の要素だがこちらは複数取得
				intSlice = append(intSlice, A[k])
			}

			// fmt.Println(intSlice)

			// intSlice内の平均値を求める
			sum := 0
			for _, in := range intSlice {
				sum+=in
			}
			var avg2 float64 = float64(sum)/float64(len(intSlice))

			// fmt.Println(avg2)
			if avg2 < max {
				max = avg2
				minIndex = i
			}

		}
	}

	return minIndex
}