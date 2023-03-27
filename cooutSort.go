package main

func countingSort(arr []int, maxValue int) []int {
	bucketlen := maxValue + 1
	bucket := make([]int, bucketlen)

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}
	for j := 0; j < bucketlen; j++ {
		if bucket[j] > 0 {
			arr[sortedIndex] = j
			bucket[j] -= 1
			sortedIndex += 1
		}
	}
	return arr
}
