package codity

func Solution(A []int, K int) []int {
	if len(A) > 0 {
		for i:=0; i<K; i++ {
			back := A[len(A)-1] // 末尾の取得
			A = A[:len(A)-1] // 末尾以外を配列に更新
			A = append([]int{back}, A[0:]...) // 末尾要素を配列の先頭に入れる
		}
	}
	return A
}