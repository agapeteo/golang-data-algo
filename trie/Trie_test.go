package trie

import "testing"

func TestTrie_Contains(t1 *testing.T) {
	trie := NewTrie()
	trie.Add("cat")
	trie.Add("can")
	trie.Add("cast")
	trie.Add("boost")

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "cat", args: args{str: "cat"}, want: true},
		{name: "can", args: args{str: "can"}, want: true},
		{name: "cas", args: args{str: "cas"}, want: false},
		{name: "casting", args: args{str: "casting"}, want: false},
		{name: "cost", args: args{str: "cost"}, want: false},
		{name: "boost", args: args{str: "boost"}, want: true},
		{name: "but", args: args{str: "but"}, want: false},
		{name: "bot", args: args{str: "bot"}, want: false},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := trie.Contains(tt.args.str); got != tt.want {
				t1.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
