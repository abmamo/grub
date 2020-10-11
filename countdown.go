package main

import (
	"time"
)

// Remaining : struct to store date diff in (t can be negative)
type Remaining struct {
	t int
	d int
	h int
	m int
	s int
}

// GetRemaining : function that takes a date and returns the remaining time up to that date
func GetRemaining(t time.Time) Remaining {
	// get current time
	currentTime := time.Now()
	// get difference
	difference := t.Sub(currentTime)
	// get total
	total := int(difference.Seconds())
	// convert total to days
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)
	// return as struct
	return Remaining{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}

}
