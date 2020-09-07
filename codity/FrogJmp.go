package Codity

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	place := X
	count := 0
	condition := true

	if X == Y || Y == 0 {
		return 0
	}

	for condition {
		count++
		place+=D
		if place >= Y {
			condition = false
		}
	}

	return count
}