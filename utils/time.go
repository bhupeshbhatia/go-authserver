package utils

import "time"

type Time interface {
	AddDate(years int, months int, days int) time.Time
	String() string
}
