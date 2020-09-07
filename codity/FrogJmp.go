package Codity

// 100%
func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	if Y < X || D <= 0 {
		return 0
	}

	if (Y-X)%D == 0 {
		return (Y-X)/D
	}
	return ((Y-X)/D) + 1

}

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	place := X
	count := 0

	if X == Y || Y == 0 {
		return 0
	}

	for {
		count++
		place+=D
		if place >= Y {
			break
		}
	}

	return count
}