package build_constraints

import (
	"runtime"
	"testing"
)

func check(value string, t *testing.T) {
	var want string
	if runtime.GOOS == "linux" {
		want = "linux"
	} else {
		want = "unknown"
	}
	if value != want {
		t.Errorf("got %s; want %s", value, want)
	}
}

func TestSuffix(t *testing.T) {
	check(suffix, t)
}

func TestTag(t *testing.T) {
	check(tag, t)
}

func asm() int

func TestAsm(t *testing.T) {
	got := asm()
	var want int
	if runtime.GOOS == "linux" {
		want = 12
	} else {
		want = 34
	}
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestCgo(t *testing.T) {
	check(cgo, t)
}
