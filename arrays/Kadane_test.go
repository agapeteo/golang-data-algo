package arrays

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{s: []int{-1, 2, 3, -7, 3, 4}}, 7},
		{"test 2", args{s: []int{34, -50, 42, 14, -5, 86}}, 137},
		{"test 2", args{s: []int{34, -50, 42, 14, -5, 86}}, 137},
		{"all negative", args{s: []int{-3, -2, -1, -4}}, 0},
		{"all positive", args{s: []int{1, 1, 0, 1, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.s); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
