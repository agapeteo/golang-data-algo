package bits

import (
	"strconv"
	"testing"
)

func TestTwoInBinary(t *testing.T) {
	println(TwoInBinary())
}

func TestIntToBinaryStr(t *testing.T) {
	println(IntToBinaryStr(7))
}

func TestIntToBinaryStrconv(t *testing.T) {
	println(IntToBinaryStrconv(7))
}

func TestCheckBitSet(t *testing.T) {
	n, _ := strconv.ParseInt("10", 2, 8)
	print(CheckBitSet(1, int(n)))
}

func TestSetBit(t *testing.T) {
	n, _ := strconv.ParseInt("10", 2, 8)
	print(SetBit(0, int(n)))
}

func TestClearBit(t *testing.T) {
	n, _ := strconv.ParseInt("111", 2, 8)
	print(ClearBit(1, int(n)))
}

func Test_toggleBit(t *testing.T) {
	n, _ := strconv.ParseInt("101", 2, 8)
	print(toggleBit(1, int(n)))
}

func TestIsEven(t *testing.T) {
	print(IsEven(1))
	print(IsEven(3))
	print(IsEven(2))
	print(IsEven(10))
}