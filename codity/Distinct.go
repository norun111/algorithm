package Codity

func removeDuplicate(args []int) int {
	encountered := map[int]bool{}

	count := 0

	for i:=0; i<len(args); i++ {
		if !encountered[args[i]] {
			count++
			encountered[args[i]] = true
		}
	}
	return count
}

func Solution(A []int) int {
	return removeDuplicate(A)
}