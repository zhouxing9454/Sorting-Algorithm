package main

import (
	"fmt"
	"sort"
)

const size = 10

func Bucket(arr []int) {
	//将数组装入桶中
	bucket := make(map[int][]int)
	for _, val := range arr {
		bucket[val/10] = append(bucket[val/10], val)
	}
	//各个桶内进行从小到大排序
	for _, val := range bucket {
		sort.Ints(val)
	}
	fmt.Println("桶内排序后: ", bucket)
	//由于map无序，因此需要获取map的key的有序集合
	var keys []int
	for key := range bucket {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	fmt.Println("桶的key排序后: ", keys)
	//桶内元素转入数组中
	index := 0
	for _, key := range keys {
		for _, val := range bucket[key] {
			arr[index] = val
			index++
		}
	}
	fmt.Printf("【Bucket】排序桶的个数为: %d, 最后结果: %v\n", len(bucket), arr)
}

func main() {
	arr := [size]int{23, 56, 98, 73, 28, 51, 90, 47, 9, 1}
	fmt.Println("原数组序列:  ", arr)
	Bucket(arr[0:])
}
