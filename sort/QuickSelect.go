package sort

type Comparable interface {
	Compare(comparable Comparable) int
}

func QuickSelect(topIdx int, s []Comparable, lowIdx int, hiIdx int) Comparable {
	pivotValue := s[hiIdx]
	i := lowIdx

	for j := i; j < hiIdx; j++ {
		if s[j].Compare(pivotValue) <= 0 {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[hiIdx] = s[hiIdx], s[i]

	if topIdx != i {
		if topIdx < i {
			return QuickSelect(topIdx, s, lowIdx, i-1)
		}
		return QuickSelect(topIdx, s, i+1, hiIdx)
	}
	return s[i]
}
