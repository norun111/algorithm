package Codity

//石の壁を構築します。
//壁は真っ直ぐで長さがNメートルで、厚みは一定でなければなりません。
//ただし、場所によって高さが異なります。壁の高さは、N個の正の整数の配列Hによって指定されます。
//H [I]は、壁の高さで、左端の右側からIからI + 1メートルです。
//特に、H [0]は壁の左端の高さ、H [N-1]は壁の右端の高さです。

func Solution(H []int) int {
	// write your code in Go 1.4
	block_cnt := 0
	stack := make([]int, 0)

	for _, height := range H {
		for len(stack) != 0 && stack[len(stack)-1] > height {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 && stack[len(stack)-1] == height {
			continue
		} else {
			block_cnt++
			stack = append(stack, height)
		}
	}

	return block_cnt
}