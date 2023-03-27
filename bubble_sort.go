package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 9, 7, 8, 4, 5, 6, 2, 1, 0}
	arr = Bubblesort(arr)
	fmt.Println(arr)
}

func Bubblesort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
