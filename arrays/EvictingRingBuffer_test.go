package arrays

import (
	"fmt"
	"testing"
)

func TestCreateRingBuffer_All(t *testing.T) {
	buf := NewRingBuffer(3)

	buf.Enqueue(1)
	buf.Enqueue(2)
	buf.Enqueue(3)
	buf.Enqueue(4)
	buf.Enqueue(5)

	res, _ := buf.Dequeue()
	fmt.Printf("res: %v, elements: %v", res, buf.Elements())
}
