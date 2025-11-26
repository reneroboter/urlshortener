package main

import "testing"

func TestExample(t *testing.T) {
	if 1+1 != 2 {
		t.Error("math is broken")
	}
}
