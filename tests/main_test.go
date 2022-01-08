package main

import (
	"testing"
)

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"valid-data", 100.0, 0.0, 0.0, true},
}

func TestDivide(t *testing.T) {
	for _, test := range tests {
		res, err := divide(test.divident, test.divisor)
		if err != nil && test.isErr != true {
			t.Error("Not expected behaviour")
		} else if res != test.expected {
			t.Error("Not expected result")
		}
	}
}

/*
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("should have Got an error")
	}
}
*/

// go test --coverprofile=coverage.out && go tool cover -html=coverage.out
// go test -v
// go test -cover
