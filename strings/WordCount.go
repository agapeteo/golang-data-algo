package strings

import "strings"

func WordCount(text string) map[string]int {
	words := make(map[string]int)
	for _, w := range strings.Split(text, " ") {
		_, ok := words[w]
		if ok {
			words[w] += 1
		} else {
			words[w] = 1
		}
	}
	return words
}
