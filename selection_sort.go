package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 9, 7, 8, 4, 5, 6, 2, 1, 0}
	arr = selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
