package main
import (
	"fmt"
)
func main(){
	var i, j int

	for i=1; i<=9; i++ {
		for j=1; j<=9; j++ {
			var result = i*j
			if result < 10 {
				fmt.Printf("  %d", result)
			} else {
				fmt.Printf(" %d", result)
			}
		}
		fmt.Println("\n")
	}
}