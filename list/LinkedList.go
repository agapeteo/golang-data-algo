package list

type LinkedList interface {
	// O(1)
	AddFirst(value interface{})

	// O(1)
	AddLast(value interface{})

	// O(1)
	Append(value interface{})

	// O(1)
	DeleteFirst()

	// O(1)
	DeleteLast()

	// O(n)
	// return copy of list
	Elements() []interface{}

	// O(n)
	// return copy of list in reversed order
	ReverseElements() []interface{}

	// O(n)
	Get(index int) (value interface{}, error error)

	// O(n)
	FirstIndexOf(element interface{}) int

	// O(n)
	InsertAfterIndex(index int, element interface{}) error

	// O(1)
	GetFirst() interface{}

	// O(1)
	GetLast() interface{}

	// O(1)
	Size() int

}
