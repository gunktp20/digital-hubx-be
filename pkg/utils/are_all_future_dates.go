package utils

import (
	"fmt"
	"time"
)

func AreAllFutureDates(dates []time.Time) (bool, error) {
	for _, date := range dates {
		isFuture, err := IsFutureDate(date)
		if err != nil {
			return false, fmt.Errorf("error for date '%s': %v", date, err)
		}
		if !isFuture {
			return false, fmt.Errorf("date '%s' is not a future date", date)
		}
	}
	// ถ้าทุกวันเป็นวันในอนาคต
	return true, nil
}
