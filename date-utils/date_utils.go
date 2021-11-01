package date_utils

import "time"

const (
	timeFormat   = "2012-01-02T15:04:05Z"
	timeDBFormat = "2006-01-02 15:04:05"
)

func TimeNow() string {
	return GetNow().Format(timeFormat)
}

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowDbFormat() string {
	return GetNow().Format(timeDBFormat)
}
