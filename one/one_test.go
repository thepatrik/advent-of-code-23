package one

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	answer := PartOne("./data")
	if answer != 55002 {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo("./data")
	if answer != 55093 {
		t.Fail()
	}
}
