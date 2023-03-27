package main

func shellSort(array []int, length int) {
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ {
			var j = i
			for {
				if j-gap < 0 || array[j] >= array[j-gap] {
					break
				}
				swap(array, j, j-gap)
				j = j - gap
			}
		}
	}
}

func swap(array []int, a int, b int) {
	array[a] = array[a] + array[b]
	array[b] = array[a] - array[b]
	array[a] = array[a] - array[b]
}
