package list

type Queue interface {
	Enqueue(element interface{})

	Dequeue() (interface{}, error)
}
