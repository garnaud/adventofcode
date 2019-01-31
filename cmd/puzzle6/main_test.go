package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	maxSize, safetySize := compute(data_test, 32)
	if maxSize != 17 || safetySize != 16 {
		fmt.Printf("maxSize= %+v\n", maxSize)
		fmt.Printf("safetySize= %+v\n", safetySize)
		t.Fail()
	}
}

var data_test = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
