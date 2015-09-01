package main

import "testing"

//here we define the input and expected output

var mul_tests = []struct {
	a, b     int
	expected int
}{
	{1, 1, 1},
	{2, 2, 4},
	{3, 3, 9},
	{4, 4, 16},
	{5, 5, 25},
}

func TestMul(t *testing.T) {
	for _, mt := range mul_tests {
		if v := Mul(mt.a, mt.b); v != mt.expected {
			t.Errorf("Mul(%d, %d) return %d, expected %d", mt.a, mt.b, v, mt.expected)
		}
	}

}
