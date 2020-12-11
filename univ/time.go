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
		return fmt.Sprintf("约 %.0f 分钟前", diff.Minutes())
	case hours < 24.0:
		return fmt.Sprintf("约 %.0f 小时前", hours)
	case hours < 72.0:
		return fmt.Sprintf("约 %.0f 天前", hours/24.0)
	default:
		if now.Year() == t.Year() { // 同一年，不用年份
			return t.Format("01-02 15:04")
		}
		return t.Format("2006-01-02")
	}
}
