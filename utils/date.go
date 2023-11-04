package utils

import (
	"time"
)

func GetBeginningOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
}

func FormatBeginningOfMonth(date time.Time) string {
	return date.Format("2006-01-02")
}
