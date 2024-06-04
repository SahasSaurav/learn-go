package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 3
	sum := Add(1, 2)
	t.Logf("%v", sum)
	if sum != expected {
		t.Errorf("result is %d instead of %d", sum, expected)
	}
}
