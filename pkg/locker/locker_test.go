package locker

import (
	"testing"
)

func TestLockerStructDefaults(t *testing.T) {

	l := NewLocker("", []string{"echo", "hi"})

	if l.PID == 0 {
		t.Errorf("Locker run failed to obtain a PID")
		t.Fail()
	}
}
