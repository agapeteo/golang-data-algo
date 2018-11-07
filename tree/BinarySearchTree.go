package tree

type BinarySearchTree interface {

	// O(log n)
	Add(newValue int)

	// O(N)
	OrderedElements() []int

	// O(log N)
	Contains(value int) bool

	// O(N)
	Depth() int

}
