package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func fizzBuzz(n int32) {
	// Write your code here

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	fizzBuzz(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
//2
func reverse(str string) string {
	compareStr := str
	li := []rune(str)
	for i, j := 0, len(str)-1; i<j; i, j = i+1, j-1 {
		li[i], li[j] = li[j], li[i]
	}

	fmt.Println(compareStr)
	if compareStr == string(li) {
		fmt.Println("palindrome")
	}
	return str
}

func breakPalindrome(palindromeStr string) string {
	// Write your code here
	reverse(palindromeStr)
	fmt.Println(palindromeStr)
	return palindromeStr
}


