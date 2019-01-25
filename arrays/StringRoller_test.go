package arrays

import "testing"

const example = "ABCD"

type input struct {
	str   string
	count int
}

type testCase struct {
	in       input
	expected string
}

func TestRollRight(t *testing.T) {
	cases := []testCase{
		{input{example, 1}, "DABC"},
		{input{example, 2}, "CDAB"},
		{input{example, len([]rune(example))}, "ABCD"},
	}

	for _, eachCase := range cases {
		actual := RollRight(eachCase.in.str, eachCase.in.count)

		if actual != eachCase.expected {
			t.Errorf("expected %v but got %v", eachCase.expected, actual)
		}
	}
}

func TestRollLeft(t *testing.T) {
	cases := []testCase{
		{input{example, 1}, "BCDA"},
		{input{example, 2}, "CDAB"},
		{input{example, len([]rune(example))}, "ABCD"},
	}

	for _, eachCase := range cases {
		actual := RollLeft(eachCase.in.str, eachCase.in.count)

		if actual != eachCase.expected {
			t.Errorf("expected %v but got %v", eachCase.expected, actual)
		}
	}
}
