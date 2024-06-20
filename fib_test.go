package framework

import (
	"testing"
)

func TestFib(t *testing.T) {
	got := Fib(3)
	want := 2
	if got != want {
		t.Errorf("expect %d, but get %d", want, got)
	} else {
		t.Logf("fib(%d)=%d", 3, got)
	}
}
