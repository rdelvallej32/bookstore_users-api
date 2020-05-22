package date_util

import "time"

const (
	apiDateFormat = "2006-01-02T15:04:05Z"
	apiDBFormat   = "2006-01-02T15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBFormat)
}
