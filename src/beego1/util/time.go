package util

import (
	"time"
)

func Dawn() *time.Time {
	now := time.Now()
	t := now.Round(24 * time.Hour)
	if t.After(now) {
		t = t.AddDate(0, 0, -1)
	}
	return &t
}
