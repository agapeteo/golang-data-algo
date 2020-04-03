package list

import "testing"

func Test_validateMultiple(t *testing.T) {
	open := []rune{'(', '{', '['}
	close := []rune{')', '}', ']'}

	type args struct {
		open  []rune
		close []rune
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should be valid 1", args{open, close, "to ((be)) or (not) to be"}, true},
		{"should be valid 2", args{open, close, "to ([be]) or {not} to be"}, true},
		{"should not valid", args{open, close, "to ( [be] {or} not }"}, false},
		{"should not valid, starts with close ", args{open, close, ")("}, false},
		{"should not valid, not closing properly ", args{open, close, "to (be) }{ or not"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateMultiple(tt.args.open, tt.args.close, tt.args.str); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}