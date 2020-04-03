package bits

import (
	"fmt"
	"strconv"
)

func TwoInBinary() int64 {
	base := 2
	bitSize := 8

	n, _ := strconv.ParseInt("10", base, bitSize)
	return n
}

func IntToBinaryStr(number int64) string {
	return fmt.Sprintf("%b", number)
}

func IntToBinaryStrconv(number int64) string {
	return strconv.FormatInt(number, 2)
}

func CheckBitSet(idx uint, n int) bool {
	mask := 1 << idx
	return (n & mask) != 0
}

func SetBit(idx uint, n int) int {
	mask := 1 << idx
	return n | mask
}

func ClearBit(idx uint, n int) int {
	mask := ^(1 << idx)
	return n & mask
}

func toggleBit(idx uint, n int) int {
	mask := 1 << idx
	return n ^ mask
}

func IsEven(n int) bool {
	return (1 & n) == 0
}
