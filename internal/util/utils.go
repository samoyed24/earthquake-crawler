package util

import (
	"time"
)

func GetCurrentJapanTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	tokyoTime := time.Now().In(loc)
	return tokyoTime
}
