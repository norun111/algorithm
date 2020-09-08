package Codity

// 100%
func Solution(N int) int {
	cnt := 0
	i := 1

	for i*i <= N {
		if N%i==0 {
			if i*i == N {
				cnt++
			} else {
				cnt+=2
			}
		}
		i++
	}

	return cnt
}

// 2147,395,600 runtime error

func Solution(N int) int {
	// write your code in Go 1.4
	count := 0
	for i:=1; i<N+1; i++ {

		if N%i==0 {
			count++
		}
	}

	return count
}