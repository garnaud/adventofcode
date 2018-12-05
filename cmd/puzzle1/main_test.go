package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func Test_compute1(t *testing.T) {
	r := compute(data1)
	if r != 2 {
		t.Fail()
	}
}

func Test_compute2(t *testing.T) {
	r := compute(data2)
	if r != 10 {
		t.Fail()
	}
}

func Test_compute3(t *testing.T) {
	r := compute(data3)
	if r != 5 {
		t.Fail()
	}
}

func Test_compute4(t *testing.T) {
	r := compute(data4)
	if r != 14 {
		t.Fail()
	}
}

var data1 = `+1
-2
+3
+1`

var data2 = `+3
+3
+4
-2
-4`

var data3 = `-6
+3
+8
+5
-6`

var data4 = `+7
+7
-2
-7
-4`
