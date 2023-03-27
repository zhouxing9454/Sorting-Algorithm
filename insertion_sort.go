package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 9, 7, 8, 4, 5, 6, 2, 1, 0}
	arr = insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}
