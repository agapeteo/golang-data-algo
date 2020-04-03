package list

func validateMultiple(open []rune, close []rune, str string) bool {
	openChars := mapFor(open)
	closeChars := mapFor(close)
	var stack []rune

	runes := []rune(str)
	for _, c := range runes {
		_, isOpen := openChars[c]
		_, isClose := closeChars[c]

		if isOpen {
			stack = append(stack, c) // push
		} else if isClose {
			if len(stack) == 0 {
				return false
			}
			lastChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			if closeChars[c] != openChars[lastChar] {
				return false
			}
		}
	}
	return len(stack) == 0
}

func mapFor(chars []rune) map[rune]int {
	m := make(map[rune]int)
	for i, c := range chars {
		m[c] = i
	}
	return m
}
