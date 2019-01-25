package arrays

func RollRight(in string, count int) string {
	runes := []rune(in)
	size := len(runes)

	for c := 0; c < count; c++ {
		var last = runes[size-1]
		for i := size - 1; i > 0; i-- {
			runes[i] = runes[i-1]
		}
		runes[0] = last
	}
	return string(runes)
}

func RollLeft(in string, count int) string {
	runes := []rune(in)
	size := len(runes)

	for c := 0; c < count; c++ {
		var first = runes[0]
		for i := 0; i < size-1; i++ {
			runes[i] = runes[i+1]
		}
		runes[size-1] = first
	}
	return string(runes)
}
