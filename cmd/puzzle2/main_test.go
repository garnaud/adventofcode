package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func Test_compute1(t *testing.T) {
	r := compute(dataTest)
	if r != 12 {
		t.Fail()
	}
}

var dataTest = `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
