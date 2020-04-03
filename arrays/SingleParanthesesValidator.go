package arrays

func validate(open rune, close rune, str string) bool {
	counter := 0
	runes := []rune(str)
	for _, rune := range runes {
		if rune == open {
			counter++
		} else if rune == close {
			if counter == 0 {
				return false
			}
			counter--
		}
	}
	return counter == 0
}
