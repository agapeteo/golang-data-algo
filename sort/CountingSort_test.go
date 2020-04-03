package sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	s := []int{3, 4, 3, 1, 12}
	CountingSort(s)
	fmt.Printf("%v", s)
}