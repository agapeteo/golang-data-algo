package math

func Primes(lim int) []int {
	s := make([]bool, lim)

	for i := 2; i*i < lim; i++ {
		if s[i] {
			continue
		}
		for j := i + i; j < lim; j += i {
			s[j] = true // true == number is not prime
		}
	}
	return extractPrimes(s)
}

func extractPrimes(s []bool) []int {
	res := make([]int, 0)
	for num, notPrime := range s {
		if num == 0 || num == 1 {
			continue
		}
		if !notPrime {
			res = append(res, num)
		}
	}
	return res
}
