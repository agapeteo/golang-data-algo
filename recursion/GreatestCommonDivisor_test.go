package recursion

import "testing"

func TestGcd(t *testing.T) {
	actual := Gcd(8, 12)
	if actual != 4 {
		t.Errorf("expecting 4, but got %d", actual)
	}
}
