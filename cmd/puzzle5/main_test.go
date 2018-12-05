package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	result := compute(data, false, ' ')
	if result != 10 {
		fmt.Printf("result = %+v\n", result)
		t.Fail()
	}
}

var data = `dabAcCaCBAcCcaDA`
