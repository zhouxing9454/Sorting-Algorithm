package main

import (
	"fmt"
	"math"
)

func radixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// 计算最大值，并确定要进行几轮基数排序
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	rounds := digitCount(maxVal)

	// 依次对每一位进行基数排序
	for i := 0; i < rounds; i++ {
		buckets := make([][]int, 10)
		for _, num := range arr {
			digit := (num / int(math.Pow(10, float64(i)))) % 10 //这段代码的作用是从一个数 num 中提取它的第 i 位数字
			buckets[digit] = append(buckets[digit], num)
		}
		arr = make([]int, 0, len(arr))
		for j := 0; j < 10; j++ {
			arr = append(arr, buckets[j]...)
		}
		fmt.Println(arr)
	}

	return arr
}

func digitCount(num int) int {
	if num == 0 {
		return 1
	}
	count := 0
	for num != 0 {
		count++
		num /= 10
	}
	return count
}

func main() {
	arr := []int{23, 56, 98, 73, 280, 51, 90, 47, 9, 1}
	fmt.Println("原数组序列:  ", arr)
	fmt.Println(radixSort(arr))
}
