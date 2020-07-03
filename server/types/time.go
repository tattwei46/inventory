package types

import (
	"fmt"
	"time"
)

const YYYYMMDD_hhmmss = "2006-01-02"
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

type Range struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (r Range) ToUnix() RangeUnix {

	loc := Location(TimeZone, time.UTC)

	from, err := time.ParseInLocation(YYYYMMDD_hhmmss, r.From, loc)
	if err != nil {
		return RangeUnix{}
	}

	tempTo, err := time.ParseInLocation(YYYYMMDD_hhmmss, r.To, loc)
	if err != nil {
		return RangeUnix{}
	}

	fmt.Println(tempTo)
	// Front end is sending range from 2020/1/1 to 2020/1/2 so we need to include 2020/1/2 as well
	to := tempTo.Add(24 * time.Hour)
	fmt.Println(to)

	return RangeUnix{
		From: from.Unix(),
		To:   to.Unix(),
	}
}

func ToUnix(created string) int64 {

	loc := Location(TimeZone, time.UTC)

	c, err := time.ParseInLocation(YYYYMMDD_hhmmss, created, loc)
	if err != nil {
		return 0
	}

	return c.Unix()
}

type RangeUnix struct {
	From int64
	To   int64
}

func (r RangeUnix) IsEmpty() bool {
	return r.To == 0 && r.From == 0
}
