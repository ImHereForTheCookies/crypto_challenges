package set1

import (
	"testing"
)

func TestChal1(t *testing.T) {
	ans := chal1("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if ans != expected {
		t.Errorf("Challenge returned %s. Wanted %s", ans, expected)
	}
}