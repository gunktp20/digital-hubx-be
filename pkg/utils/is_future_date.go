package utils

import (
	"errors"
	"time"
)

func IsFutureDate(date time.Time) (bool, error) {

	// ดึงวันที่ปัจจุบัน (แค่ปี, เดือน, วัน)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// ตรวจสอบว่าวันที่เป็นในอดีตหรือวันนี้
	if date.Before(today) || date.Equal(today) {
		return false, errors.New("the event date cannot be today or in the past")
	}

	// เป็นวันที่ในอนาคต
	return true, nil
}
