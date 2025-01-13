package utils

import (
	"time"
)

func IsEventDateValidForReg(eventDate time.Time) (bool, error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if today.Equal(eventDate) || today.After(eventDate) {
		return false, nil
	}

	return true, nil
}
