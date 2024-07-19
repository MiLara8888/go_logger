package syslog

import "testing"

func TestSysLog(t *testing.T) {

	l, err := New(func(o *Options) {
	})
	if err != nil {
		t.Error(err)
		return
	}

	l.Criticalf("%s %d", "test", 1)
}
