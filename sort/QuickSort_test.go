package sort

import "testing"

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		exp  []int
	}{
		{"ascending", args{[]int{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
		{"descending", args{[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"tricky middle", args{[]int{4, 2, 0, 5, 1}}, []int{0, 1, 2, 4, 5}},
		{"simple example with dups", args{[]int{1, 9, 3, 5, 4, 1}}, []int{1, 1, 3, 4, 5, 9}},
		{"mirrored dups", args{[]int{5, 5, 3, 5, 7, 5, 5}}, []int{3, 5, 5, 5, 5, 5, 7}},
		{"mirrored dups 2", args{[]int{5, 5, 7, 8, 5, 3, 5, 1, 2, 3}}, []int{1, 2, 3, 3, 5, 5, 5, 5, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PickedQuickSortStrategy = Lomuto
			QuickSort(tt.args.arr)
			if !isEqual(tt.args.arr, tt.exp) {
				t.Errorf("test `%v` failed. expecteed %v but got %v. Strategy # %v", tt.name, tt.exp, tt.args.arr, PickedQuickSortStrategy)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			PickedQuickSortStrategy = Hoare
			QuickSort(tt.args.arr)
			if !isEqual(tt.args.arr, tt.exp) {
				t.Errorf("test `%v` failed. expecteed %v but got %v. Strategy # %v", tt.name, tt.exp, tt.args.arr, PickedQuickSortStrategy)
			}
		})
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
