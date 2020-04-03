package list

type MinStack struct {
	stack    []int
	minStack []int
}

func CreateMinStack() *MinStack {
	stack := new(MinStack)
	stack.stack = make([]int, 0)
	stack.minStack = make([]int, 0)

	return stack
}

func (s *MinStack) Push(value int) {
	if len(s.minStack) == 0 || value <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, value)
	}
	s.stack = append(s.stack, value)
}

func (s *MinStack) Pop() int {
	result := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	if result == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
	return result
}

func (s *MinStack) Min() int {
	return s.minStack[len(s.minStack)-1]
}
