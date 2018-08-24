package algorithm

import (
	"testing"
)

func TestHuiwen(t *testing.T) {
	if Huiwen(123) {
		t.Log("123 是回文")
	}
	if Huiwen(121) {
		t.Log("121 是回文")
	}
	if Huiwen(11) {
		t.Log("11 是回文")
	}
	if Huiwen(999) {
		t.Log("999 是回文")
	}
}

func TestChecken(t *testing.T) {
	Checken(100, 100)
}
func TestPeach(t *testing.T) {
	t.Log("桃子数: ", Peach(10))
}
func TestHanoi(t *testing.T) {
	Hanoi(3, 'A', 'B', 'C')
}
