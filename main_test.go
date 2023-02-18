package main_test

import (
	"testing"
)

func TestFoo(t *testing.T) {
	if "" != "" { //nolint
		t.Fatalf("fail!")
	}
}
