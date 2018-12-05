package main

import "testing"

func TestData1(t *testing.T) {
	r := compute(data1)
	if r != 4 {
		t.Fail()
	}
}

var data1 = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
