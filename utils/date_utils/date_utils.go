package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T05:04:20.123Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}