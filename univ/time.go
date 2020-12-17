package univ

import (
	"fmt"
	"time"
)

// TimeAgo see code
func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	hours := diff.Hours()
	switch {
	case hours < 1.0:
		return fmt.Sprintf("%.0f minute ago", diff.Minutes())
	case hours < 24.0:
		return fmt.Sprintf("%.0f hour ago", hours)
	case hours < 72.0:
		return fmt.Sprintf("%.0f day ago", hours/24.0)
	default:
		if now.Year() == t.Year() { // 同一年，不用年份
			return t.Format("01-02 15:04")
		}
		return t.Format("2006-01-02")
	}
}
