package sort

import (
	"reflect"
	"testing"
)

type comparableInt int

func (i comparableInt) Compare(other Comparable) int {
	return int(i) - int(other.(comparableInt))
}

func TestQuickSelect(t *testing.T) {
	s := []Comparable{comparableInt(2), comparableInt(1), comparableInt(4), comparableInt(9), comparableInt(7)}

	type args struct {
		topIdx int
		s      []Comparable
		lowIdx int
		hiIdx  int
	}
	tests := []struct {
		name string
		args args
		want Comparable
	}{
		{"0", args{0, s, 0, len(s) - 1}, comparableInt(1)},
		{"1", args{1, s, 0, len(s) - 1}, comparableInt(2)},
		{"2", args{2, s, 0, len(s) - 1}, comparableInt(4)},
		{"3", args{3, s, 0, len(s) - 1}, comparableInt(7)},
		{"4", args{4, s, 0, len(s) - 1}, comparableInt(9)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSelect(tt.args.topIdx, tt.args.s, tt.args.lowIdx, tt.args.hiIdx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
