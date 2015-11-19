package timex

import (
	"time"
)

var UTC8 = time.FixedZone("UTC+8", 8*60*60)

const (
	SecondsPerMinute = 60
	SecondsPerHour   = 60 * 60
	SecondsPerDay    = 24 * SecondsPerHour
	SecondsPerWeek   = 7 * SecondsPerDay
	DaysPer400Years  = 365*400 + 97
	DaysPer100Years  = 365*100 + 24
	DaysPer4Years    = 365*4 + 1
)

func Now() int64 {
	return ToUnix(time.Now())
}

func ToUnix(t time.Time) int64 {
	return t.In(UTC8).Unix()
}

func FromUnix(i int64) time.Time {
	return time.Unix(i, 0).In(UTC8)
}

func Today0() time.Time {
	year , month , day  := time.Now().In(UTC8).Date()
	t := time.Date(year, month, day, 0, 0, 0, 0 , UTC8)
	return t
}

func NextDay0() time.Time {
	year , month , day  := time.Now().In(UTC8).Date()
	t := time.Date(year, month, day+1, 0, 0, 0, 0 , UTC8)
	return t
}
