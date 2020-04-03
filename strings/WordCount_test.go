package strings

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	text := "long long long way to to my home"
	words := WordCount(text)
	if words["long"] != 3 || words["way"] != 1 || words["to"] != 2 || words["my"] != 1 || words["home"] != 1 {
		t.Errorf("wrong results, words: %v", words)
	}
}