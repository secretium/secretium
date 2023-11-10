package helpers

import "time"

// DatetimeChecker compares two datetimes and returns true if they are equal or a is after b.
func DatetimeChecker(a, b time.Time) bool {
	return a.Equal(b) || a.After(b)
}
