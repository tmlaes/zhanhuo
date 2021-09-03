package utils

import "time"

const DATE_FORMAT = "2006-01-02 15:04:05"

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}
