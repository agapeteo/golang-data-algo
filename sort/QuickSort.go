package sort

type QuickSortingStrategy int

const (
	Lomuto QuickSortingStrategy = iota
	Hoare
)

var PickedQuickSortStrategy = Lomuto

func QuickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	quickSortRange(arr, 0, len(arr)-1)
}

func quickSortRange(arr []int, lowIdx, highIdx int) {
	if lowIdx >= highIdx {
		return
	}

	var partitionIdx = -1

	if PickedQuickSortStrategy == Lomuto {
		partitionIdx = partitionLomuto(arr, lowIdx, highIdx)
	}
	if PickedQuickSortStrategy == Hoare {
		partitionIdx = partitionHoare(arr, lowIdx, highIdx)
	}
	if partitionIdx == -1 {
		panic("PickedQuickSortStrategy not defined")
	}

	quickSortRange(arr, lowIdx, partitionIdx-1)
	quickSortRange(arr, partitionIdx+1, highIdx)
}

func partitionLomuto(arr []int, lowIdx, highIdx int) int {
	pivot := arr[highIdx]
	i := lowIdx

	for j := lowIdx; j < highIdx; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[highIdx], arr[i] = arr[i], arr[highIdx]
	return i
}

func partitionHoare(arr []int, lowIdx, highIdx int) int {
	pivot := arr[lowIdx]
	leftIdx := lowIdx
	rightIdx := highIdx + 1

	for {
		for {
			leftIdx++
			if leftIdx == highIdx || arr[leftIdx] >= pivot {
				break
			}
		}
		for {
			rightIdx--
			if  rightIdx == lowIdx || arr[rightIdx] <= pivot{
				break
			}
		}
		if leftIdx >= rightIdx {
			break
		}
		arr[leftIdx], arr[rightIdx] = arr[rightIdx], arr[leftIdx]
	}
	arr[rightIdx], arr[lowIdx] = arr[lowIdx], arr[rightIdx]
	return rightIdx
}
