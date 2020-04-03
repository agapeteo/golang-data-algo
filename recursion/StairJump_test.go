package recursion

import "testing"

func TestStairJumpCount(t *testing.T) {
	ways := []int{1, 2, 5}
	type args struct {
		stairsCount int
		jumps       []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, ways}, 1},
		{"2", args{2, ways}, 2},
		{"3", args{3, ways}, 3},
		{"4", args{4, ways}, 5},
		{"5", args{5, ways}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StairJumpCount(tt.args.stairsCount, tt.args.jumps); got != tt.want {
				t.Errorf("StairJumpCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
