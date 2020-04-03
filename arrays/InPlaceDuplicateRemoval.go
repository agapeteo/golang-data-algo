package arrays

import "sort"

func RemoveDuplicates(arr []string) int {
	sort.Strings(arr)
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[i] {
			arr[i+1] = arr[j]
			i++
		}
	}
	return i
}
