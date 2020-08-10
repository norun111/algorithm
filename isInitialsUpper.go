package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isInitialsUpper(hoge string) bool {
	if hoge == "" {
		return false
	}
	// 文字列の一文字目をrune型にキャストしています。
	initials := rune(hoge[0])

	// unicodeパッケージのIsUpper関数を使用しています。
	// 引数にrune型の変数initialsをいれ、
	// そのrune文字が大文字かどうかをbool値で返します。
	return unicode.IsUpper(initials)
}

func main(){
	// Your code here!
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	hoge := sc.Text()
	fmt.Println(isInitialsUpper(hoge))
}