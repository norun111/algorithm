package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intSlices, _ = SplitSlice([]int{1, 2, 3, 4, 5}, 2)
	for _, is := range intSlices {
		fmt.Printf("int slice %v \n", is.([]int))
	}

	var strSlices, _ = SplitSlice([]string{"A", "B", "C", "D", "E"}, 2)
	for _, ss := range strSlices {
		fmt.Printf("string slice %v \n", ss.([]string))
	}
}

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

