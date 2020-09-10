package main

func index(slice []string, str string) int {

	index := -1
	for i, s := range slice {

		if s == str {
			index = i
			break //一番初めだけ検索する
		}
	}

	return index // indexが-1は配列に文字列が存在しない場合
}

//単語が配列にあるかどうかの判断関数
func isContain(slice []string, str string) bool {

	condition := false
	for _, s := range slice {

		if s == str {
			condition = true
			break //一番初めだけ検索する
		}
	}

	return condition //なければfalseを返す
}

//単語が配列にあるかどうかの判断関数
func isContain(slice []int, str int) bool {

	condition := false
	for _, s := range slice {

		if s == str {
			condition = true
			break //一番初めだけ検索する
		}
	}

	return condition //なければfalseを返す
}
