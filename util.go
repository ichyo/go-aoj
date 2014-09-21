package aoj

import "time"

func UnixToTime(unix uint64) time.Time {
	return time.Unix(int64(unix/1000), int64(unix%1000)*1000000)
}
