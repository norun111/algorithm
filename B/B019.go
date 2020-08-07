package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"reflect"
)

func SplitSlice(list interface{}, size int) ([]interface{}, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size need positive number")
	}

	// sliceかどうかのチェック
	if reflect.TypeOf(list).Kind() != reflect.Slice {
		return nil, fmt.Errorf("src must be slice")
	}

	var result = []interface{}{}

	var v = reflect.ValueOf(list)
	for i := 0; i < v.Len(); i += size {
		if i+size <= v.Len() {
			result = append(result, v.Slice(i, i+size).Interface())
		} else {
			result = append(result, v.Slice(i, v.Len()).Interface())
		}
	}

	return result, nil
}

func main(){

	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	l := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(l[0])
	K, _ := strconv.Atoi(l[1])

	blockNum := N/K * N/K

	blockList := make([][]int, blockNum)

	// fmt.Println(blockList)

	num := N/K
	blockIndex := 0

	for i:=0; i<N; i++ {

		sc.Scan()
		list := strings.Split(sc.Text(), " ")

		iIndex := i + 1

		for j:=0; j<len(list); j++ {
			// fmt.Println(blockIndex)
			pixel, _ := strconv.Atoi(list[j])
			blockList[blockIndex] = append(blockList[blockIndex], pixel)

			in := j + 1
			if in%K == 0{
				if j == len(list)-1 {
					if iIndex%K == 0 {
						blockIndex+=num
					}
					blockIndex-=num
				}
				blockIndex++
			}
		}

	}
	// fmt.Println(blockList)

	resultList := []int{}

	for _, b := range blockList {

		sum := 0

		for _, b2 := range b {
			sum+=b2
		}
		length := len(b)
		result := sum/length

		// fmt.Println(result)

		resultList = append(resultList, result)
	}

	strList := []string{}

	for _, r := range resultList {
		str := strconv.Itoa(r)
		strList = append(strList, str)
	}


	var intSlice, _ = SplitSlice(strList, num)
	for _, is := range intSlice {
		fmt.Println(strings.Join(is.([]string), " "))
	}
}