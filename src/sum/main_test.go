package main

import (
	"testing"
)

func TestSun(t *testing.T) {
	got := sum(5, 5)
	wanted := 10

	if got != wanted {
		t.Errorf("sum(5, 5) \n got: %v \n wanted: %v", got, wanted)
	}
}
