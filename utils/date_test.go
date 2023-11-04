package utils

import (
	"testing"
	"time"
)

func TestGetBeginningOfMonth(t *testing.T) {
	date := GetBeginningOfMonth()
	if date.Day() != 1 {
		t.Fatalf("Expected day to be 1, got %d", date.Day())
	}
	if date.Hour() != 0 {
		t.Fatalf("Expected hour to be 0, got %d", date.Hour())
	}
	if date.Minute() != 0 {
		t.Fatalf("Expected minute to be 0, got %d", date.Minute())
	}
	if date.Second() != 0 {
		t.Fatalf("Expected second to be 0, got %d", date.Second())
	}
	if date.Nanosecond() != 0 {
		t.Fatalf("Expected nanosecond to be 0, got %d", date.Nanosecond())
	}
}

func TestFormatBeginningOfMonth(t *testing.T) {
	date := FormatBeginningOfMonth(time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC))
	if date != "2021-10-01" {
		t.Fatalf("Expected date to be 2021-10-01, got %s", date)
	}
}
