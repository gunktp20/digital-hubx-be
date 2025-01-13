package utils

import (
	"time"
)

// ฟังก์ชันตรวจสอบว่าเป็น ISO 8601 หรือไม่
func IsValidISO8601(dateStr string) bool {
	// ใช้ time.Parse เพื่อแปลงจาก string เป็น time.Time ตามรูปแบบ RFC3339 (ISO 8601)
	_, err := time.Parse(time.RFC3339, dateStr)
	return err == nil
}
