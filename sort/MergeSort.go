package sort

func MergeSort(inputSlice []int) {
	if len(inputSlice) < 2 {
		return
	}
	tmpSlice := make([]int, len(inputSlice), len(inputSlice))
	mergeSortRange(inputSlice, tmpSlice, 0, len(inputSlice)-1)
}

func mergeSortRange(inputSlice, tmpSlice []int, lowIdx, highIdx int) {
	if lowIdx >= highIdx {
		return
	}
	midIdx := lowIdx + (highIdx-lowIdx)/2
	mergeSortRange(inputSlice, tmpSlice, lowIdx, midIdx)
	mergeSortRange(inputSlice, tmpSlice, midIdx+1, highIdx)

	merge(inputSlice, tmpSlice, lowIdx, midIdx, highIdx)
}

func merge(inputSlice, tmpSlice []int, lowIdx, midIdx, highIdx int) {
	for i := lowIdx; i <= highIdx; i++ {
		tmpSlice[i] = inputSlice[i]
	}

	leftIdx := lowIdx
	rightIdx := midIdx + 1
	curIdx := leftIdx

	for leftIdx <= midIdx && rightIdx <= highIdx {
		if tmpSlice[leftIdx] <= tmpSlice[rightIdx] {
			inputSlice[curIdx] = tmpSlice[leftIdx]
			leftIdx++
		} else {
			inputSlice[curIdx] = tmpSlice[rightIdx]
			rightIdx++
		}
		curIdx++
	}

	remaining := midIdx - leftIdx
	for i := 0; i <= remaining; i++ {
		inputSlice[curIdx+i] = tmpSlice[leftIdx+i]
	}
}
