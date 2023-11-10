package helpers

import (
	"errors"
	"time"

	"github.com/secretium/secretium/internal/messages"
)

// ExpiresDatetimeSwitcher returns a new datetime with a duration.
func ExpiresDatetimeSwitcher(now time.Time, duration string) (time.Time, error) {
	// Parse the duration.
	switch duration {
	case "5m":
		// 5 minutes.
		return now.Add(time.Minute * 5), nil
	case "15m":
		// 15 minutes.
		return now.Add(time.Minute * 15), nil
	case "30m":
		// 30 minutes.
		return now.Add(time.Minute * 30), nil
	case "1h":
		// 1 hour.
		return now.Add(time.Hour * 1), nil
	case "3h":
		// 3 hours.
		return now.Add(time.Hour * 3), nil
	case "12h":
		// 12 hours.
		return now.Add(time.Hour * 12), nil
	case "1d":
		// 1 day.
		return now.Add(time.Hour * 24), nil
	case "3d":
		// 3 days.
		return now.Add(time.Hour * 24 * 3), nil
	case "7d":
		// 7 days.
		return now.Add(time.Hour * 24 * 7), nil
	case "14d":
		// 14 days.
		return now.Add(time.Hour * 24 * 14), nil
	case "30d":
		// 30 days.
		return now.Add(time.Hour * 24 * 30), nil
	default:
		// Error.
		return now, errors.New(messages.ErrSecretExpiresAtNotValid)
	}
}
