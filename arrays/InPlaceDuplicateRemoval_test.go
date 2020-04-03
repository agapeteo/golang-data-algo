package arrays

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []string{"a", "b", "a", "c", "b", "d"}

	idx := RemoveDuplicates(arr)
	fmt.Printf("idx %v, arr: %v", idx, arr)
}