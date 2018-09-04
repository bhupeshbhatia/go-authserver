package mocks

import (
	"time"
)

type Time struct {
	Time time.Time
}

func (t *Time) AddDate(years int, months int, days int) time.Time {
	time := t.Time.AddDate(years, months, days)
	return time
}

func (t *Time) String() string {
	return t.Time.String()
}
