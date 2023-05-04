package main

import "testing"

func TestAdd(t *testing.T) {
	if add(1, 2) == 1 {
		t.Logf("a=%v b=%v ans=%v\n", 1, 2, add(1, 2))
	}
}
