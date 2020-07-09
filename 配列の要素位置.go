package GoPaiza

func index(slice []string, str string) int {

	index := 0
	for i, s := range slice {

		if s == str {
			index = i
			break //一番初めだけ検索する
		}
	}

	return index // indexが0は配列に文字列が存在しない場合
}