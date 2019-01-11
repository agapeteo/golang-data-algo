package list

type Stack interface {
	Push(element interface{})

	Pull() (interface{}, error)

	Peek() (interface{}, error)
}
