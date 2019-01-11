package tree

type Heap interface {
	Peak() (int, error)
	Pop() (int, error)
	Push(v int)

	Elements() []int
	Size() int
}
