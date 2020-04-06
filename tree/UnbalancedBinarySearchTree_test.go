package tree

import (
	"reflect"
	"testing"
)

func TestUnbalancedBinarySearchTree_ImplementsInterface(t *testing.T) {
	// given
	// when
	var _ BinarySearchTree = &UnbalancedBinarySearchTree{}

	// then
	// should compile
}

func TestUnbalancedBinarySearchTree_Add(t *testing.T) {
	// given
	var tree = &UnbalancedBinarySearchTree{}

	// when
	tree.Add(5)
	tree.Add(3)
	tree.Add(2)
	tree.Add(1)
	tree.Add(4)
	tree.Add(10)
	tree.Add(20)
	tree.Add(7)

	// then
	expected := []int{1, 2, 3, 4, 5, 7, 10, 20}
	if !isEqual(tree.OrderedElements(), expected){
		t.Errorf("expected %v but got %v", expected, tree.OrderedElements())
	}
}

func TestUnbalancedBinarySearchTree_Contains(t *testing.T) {
	// given
	var tree = &UnbalancedBinarySearchTree{}

	// when
	tree.Add(5)
	tree.Add(3)
	tree.Add(2)
	tree.Add(1)
	tree.Add(4)
	tree.Add(10)
	tree.Add(20)
	tree.Add(7)

	// then
	if !tree.Contains(5){
		t.Errorf("tree should contain 5")
	}

	if !tree.Contains(1){
		t.Errorf("tree should contain 1")
	}

	if !tree.Contains(20){
		t.Errorf("tree should contain 20")
	}

	if tree.Contains(19){
		t.Errorf("tree should not contain 19")
	}
}

func TestUnbalancedBinarySearchTree_Contains_Empty(t *testing.T) {
	// given
	// when
	var tree = &UnbalancedBinarySearchTree{}

	// then
	if tree.Contains(19){
		t.Errorf("tree should not contain anything")
	}
}

func TestUnbalancedBinarySearchTree_TwoLevels(t *testing.T) {
	// given
	var tree = &UnbalancedBinarySearchTree{}

	// when
	tree.Add(5)
	tree.Add(3)
	tree.Add(4)
	tree.Add(6)


	// then
	if tree.Depth() != 2 {
		t.Errorf("depth should be 2, but got %v", tree.Depth())
	}
}

func TestUnbalancedBinarySearchTree_Depth_Empty(t *testing.T) {
	// given
	// when
	var tree = &UnbalancedBinarySearchTree{}

	// then
	if tree.Depth() != 0 {
		t.Errorf("depth should be 0, but got %v", tree.Depth())
	}
}

func TestUnbalancedBinarySearchTree_OnlyRoot(t *testing.T) {
	// given
	var tree = &UnbalancedBinarySearchTree{}

	// when
	tree.Add(0)

	// then
	if tree.Depth() != 0 {
		t.Errorf("depth should be 0, but got %v", tree.Depth())
	}
}

func TestUnbalancedBinarySearchTree_OneLevel(t *testing.T) {
	// given
	var tree = &UnbalancedBinarySearchTree{}

	// when
	tree.Add(0)
	tree.Add(1)


	// then
	if tree.Depth() != 1 {
		t.Errorf("depth should be 1, but got %v", tree.Depth())
	}
}

func isEqual(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestFromSorted(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int // depth
	}{
		{name: "0-10", args: args{[]int{1,2,3,4,5,6,7,8,9,10}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromSorted(tt.args.s); !reflect.DeepEqual(got.Depth(), tt.want) {
				t.Errorf("FromSorted() = %v, want %v", got.Depth(), tt.want)
			}
		})
	}
}