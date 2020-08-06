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
	// UTF-8のバイトをrune型にキャストしています
	initials := rune(hoge[0])

	// unicodeパッケージのIsUpper関数の引数にrune型の変数initialsを入れ、
	// そのrune文字が大文字かどうかをbool値で返します
	return unicode.IsUpper(initials)
}

func main(){
	// Your code here!
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	hoge := sc.Text()
	fmt.Println(isInitialsUpper(hoge))
}