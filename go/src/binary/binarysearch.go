package main

import "fmt"

func binarySearch(arr []int, srchItm int) int {
	mid := 0
	lo := 0
	hi := len(arr) - 1
	for hi > lo {
		if lo < 0 {
			return -1
		}
		mid = int((hi + lo) / 2)

		midElm := arr[mid]
		fmt.Println(lo, hi, mid, midElm)

		if midElm == srchItm {
			return mid + 1
		} else if srchItm < midElm {
			hi = mid
		} else {
			lo = mid
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binarySearch(arr, 8))

}
