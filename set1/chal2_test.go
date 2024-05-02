package set1

import (
	"testing"
)

func TestChal2(t *testing.T) {
	ans := Chal2("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	expected := "746865206b696420646f6e277420706c6179"
	if ans != expected {
		t.Errorf("Returned %s. Wanted %s", ans, expected)
	}
}