package sort

func max(s []int) int {
	max := s[0]
	for _, v := range s {
		if max < v {
			max = v
		}
	}
	return max
}

func CountingSort(s []int) {
	size := max(s) + 1
	countSlice := make([][]int, size)

	for _, v := range s {
		curSlice := &countSlice[v]
		*curSlice = append(*curSlice, v)
	}

	idx := 0
	for _, curSlice := range countSlice {
		for _, v := range curSlice {
			s[idx] = v
			idx++
		}
	}
}
