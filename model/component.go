package model

import "time"

type Failure struct {
	// down time (unplanned)
	StartTime time.Time
	// uptime (after repair)
	EndTime time.Time
	Cause   string
}

type Available struct {
	// beginning of availability
	StartTime time.Time
	// end of availability
	EndTime time.Time
}

type Component struct {
	Failures     []Failure
	Availability []Available
}
