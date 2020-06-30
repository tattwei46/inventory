package types

import "time"

const YYYYMMDD_hhmmssMST = "2006-01-02 15:04:05 MST"
const TimeZone = "Asia/Singapore"

func Format(unixSec int64, timeZone string, format string) string {
	if unixSec <= 0 {
		return ""
	}
	return time.Unix(unixSec, 0).
		In(Location(timeZone, time.UTC)).
		Format(format)
}

func Location(timeZone string, defaultValue *time.Location) *time.Location {
	loc := defaultValue
	if l, err := time.LoadLocation(timeZone); err == nil {
		loc = l
	}
	return loc
}
