package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var arrStr = flag.String("arr", "", "arr")
var val = flag.Int("val", 0, "val")

func main() {
	flag.Parse()
	arrStrArr := strings.Split(*arrStr, ",")
	arr := make([]int, len(arrStrArr))
	for idx, arg := range arrStrArr {
		item, _ := strconv.Atoi(arg)
		arr[idx] = item
	}

	fmt.Println(BinarySearch(arr, *val))
}

func BinarySearch(arr []int, val int) int {
	return binarySearch(arr, val, 0, len(arr)-1)
}

func binarySearch(arr []int, val, start, end int) int {
	arrLen := end - start + 1
	if arrLen == 1 {
		if arr[start] == val {
			return start
		} else {
			return -1
		}
	}

	if arrLen == 2 {
		if arr[start] == val {
			return start
		} else if arr[start+1] == val {
			return start + 1
		} else {
			return -1
		}
	}

	middle := arrLen/2 + start
	middleVal := arr[middle]
	if middleVal > val {
		return binarySearch(arr, val, start, middle)
	}
	if middleVal < val {
		return binarySearch(arr, val, middle, end)
	}
	if middleVal == val {
		return middle
	}

	return -1
}
