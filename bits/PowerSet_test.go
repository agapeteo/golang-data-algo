package bits

import (
	"fmt"
	"testing"
)

func TestPowerSet(t *testing.T) {
	set := make([]interface{}, 4)
	set[0] = 1
	set[1] = 2
	set[2] = 3
	set[3] = 4

	actual := PowerSet(set)

	fmt.Printf("->%v", actual)
}