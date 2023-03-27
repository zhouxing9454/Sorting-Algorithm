package main

import "fmt"

func partition(nums []int, left int, right int) int {
	value := nums[left] // 基准值
	for left < right {
		for nums[right] >= value && left < right { // 依次查找大于等于基准值的位置
			right--
		}
		nums[left] = nums[right]
		for nums[left] < value && left < right { // 依次查找小于基准值的位置
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = value
	// 最终 left == right 就是基准值的位置
	return left
}

func QuickSort(list []int, left int, right int) {
	if left < right {
		middle := partition(list, left, right)
		QuickSort(list, left, middle-1)
		QuickSort(list, middle+1, right)
	}
}

func main() {
	var arr []int
	arr = []int{1, 9, 7, 8, 4, 5, 6, 2, 1, 0}
	QuickSort(arr, 0, 9)
	fmt.Println(arr)
}
