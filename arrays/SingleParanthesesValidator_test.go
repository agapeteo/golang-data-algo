package arrays

import "testing"

func Test_validate(t *testing.T) {
	type args struct {
		open  rune
		close rune
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid case", args{rune('('), rune(')'), "to ((be)) or (not) to be"}, true},
		{"not valid case", args{rune('('), rune(')'), "to ((be)(or) not"}, false},
		{"should not start with closed", args{rune('('), rune(')'), ")("}, false},
		{"not valid closed at the end", args{rune('('), rune(')'), "to (be)) or ("}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.open, tt.args.close, tt.args.str); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}