package foo_test

import (
	"testing"
)

func TestFoo(t *testing.T) {
	if "" != "" {
		t.Fatalf("fail!")
	}
}
