package dynamic

import (
	"fmt"
	"testing"
)

func TestPowerSet(t *testing.T) {
	//type args struct {
	//	inSet []interface{}
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want [][]interface{}
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := PowerSet(tt.args.inSet); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("PowerSet() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	set := make([]interface{}, 3)
	set[0] = 1
	set[1] = 2
	set[2] = 3

	actual := PowerSet(set)

	fmt.Printf("%v", actual)
}