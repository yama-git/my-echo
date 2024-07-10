package main

import "testing"

func TestPing(t *testing.T) {
	got := ping()
	if got != "pongpang" {
		t.Errorf("ping() =%v; want pongpang", got)
	}
}
