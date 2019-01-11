package tree

import "errors"

var NoSuchElementError = errors.New("no such element")

type MaxHeap struct {
	elm []int
}

func (H *MaxHeap) Peak() (int, error) {
	if len(H.elm) == 0 {
		return -1, NoSuchElementError
	}
	return H.elm[0], nil
}

func (H *MaxHeap) Pop() (int, error) {
	if len(H.elm) == 0 {
		return -1, NoSuchElementError
	}

	var result = H.elm[0]

	if len(H.elm) == 1 {
		H.elm = make([]int, 0)
		return result, nil
	}

	H.elm[0] = H.elm[len(H.elm)-1] // put last to first
	H.elm = H.elm[:len(H.elm)-1]   // trim last element in slice

	siftDown(0, H.elm)

	return result, nil
}

func siftDown(i int, e []int) {
	if i < 0 || i >= len(e)/2 {
		return // no need to sift down leaf elements
	}

	var l = leftChildIdx(i)
	var r = rightChildIdx(i)

	var largerChild = l
	if r < len(e) && e[r] > e[l] {
		largerChild = r
	}

	if e[largerChild] > e[i] {
		swap(i, largerChild, e)
		siftDown(largerChild, e)
	}
}

func (H *MaxHeap) Push(v int) {
	if len(H.elm) == 0 {
		H.elm = append(H.elm, v)
		return
	}

	H.elm = append(H.elm, v)

	siftUp(len(H.elm)-1, H.elm)
}

func siftUp(i int, e []int) {
	if i <= 0 {
		return
	}
	var p = parentIdx(i)

	if e[i] > e[p] {
		swap(i, p, e)
		siftUp(p, e)
	}
}

func swap(i int, j int, e []int) {
	var t = e[i]
	e[i] = e[j]
	e[j] = t
}

func (H *MaxHeap) Elements() []int {
	var result = make([]int, len(H.elm))
	copy(result, H.elm)
	return result
}

func (H *MaxHeap) Size() int {
	return len(H.elm)
}

func leftChildIdx(i int) int {
	return i*2 + 1
}

func rightChildIdx(i int) int {
	return i*2 + 2
}

func parentIdx(i int) int {
	return (i - 1) / 2
}
