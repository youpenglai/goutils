package convert

import (
	"time"
)

const formatlayout = "2006-01-02 15:04:05"

func TimeToStr(t time.Time) string {
	return t.Format(formatlayout)
}

func TimeToTimestamp(t time.Time) int64 {
	return t.Unix()
}

func StrToTime(v string) (time.Time, error) {
	return time.Parse(formatlayout, v)
}

func StrToTimestamp(v string) (int64, error) {
	t, err := StrToTime(v)
	if err != nil {
		return 0, err
	}
	return TimeToTimestamp(t), nil
}

func ValueToTime(year, month, day, hour, min, sec, nsec int) (time.Time) {
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, time.Local)
}

func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func TimestampToStr(timestamp int64) string {
	return TimeToStr(TimestampToTime(timestamp))
}
