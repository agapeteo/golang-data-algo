package math

import (
	"reflect"
	"testing"
)

func TestPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"10", args{10}, []int{2, 3, 5, 7}},
		{"100", args{100}, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Primes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Primes() = %v, want %v", got, tt.want)
			}
		})
	}
}
