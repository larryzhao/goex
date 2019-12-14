package timex

import "time"

func NowUnix13() int64 {
	return ToUnix13(time.Now())
}

func ToUnix13(t time.Time) int64 {
	ns := t.UnixNano()
	return int64(ns / 1000000)
}

func ParseUnix13(timestamp int64) time.Time {
	f := float64(timestamp / 1000)
	sec := int64(f)
	nsec := int64((f - float64(sec)) * 1000000000)

	return time.Unix(sec, nsec)
}