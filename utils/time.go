package utils

import (
	"time"
)

type Time interface {
	Now() time.Time
}
