package main

import (
	"testing"
	"time"
)

func TestGetRemaining(t *testing.T) {
	// get current time
	now := time.Now()
	// get remaining
	remaining := GetRemaining(now)

	if remaining.t != 0 {
		t.Errorf("get remaining failed.")
	}
}
