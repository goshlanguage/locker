package locker

import (
	"testing"
)

func TestLockerStructDefaults(t *testing.T) {

	opts := LockerOpts{
		Name:     "testy_timothy",
		Command:  []string{"echo", "hi"},
		Hostname: "testy_timothy",
	}
	l := opts.Build()
	l.Run()

	if l.PID == 0 {
		t.Errorf("Locker run failed to obtain a PID")
		t.Fail()
	}
}
