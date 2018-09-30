package main // import "github.com/ryanhartje/locker/cmd/locker"

import (
	"testing"
)

func TestRun(t *testing.T) {
	var args []string
	cmd := newRunCmd(args)
	if cmd == nil {
		t.Errorf("Run command didn't produce anything, import error?")
	}
}
