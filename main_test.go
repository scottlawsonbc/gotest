package main

import "testing"

func TestMultiply(t *testing.T) {
	if multiply(1, 1) != 1 {
		t.Error("wrong")
	}
}
