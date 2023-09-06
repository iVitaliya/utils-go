package date

import (
	"time"

	"github.com/iVitaliya/utils-go/utils"
)

type timeInstance struct {
	millisecond int64
	second      int
	minute      int
	hour        int
}

func getInstance(time_arg time.Time) *timeInstance {
	var (
		now = time_arg.Local()
		ms  = now.UnixMilli()
		s   = now.Second()
		m   = now.Minute()
		h   = now.Hour()
	)

	return &timeInstance{
		millisecond: ms,
		second:      s,
		minute:      m,
		hour:        h,
	}
}

// This will give the current time in AM/PM
func CurrentATime() string {
	var (
		t            = getInstance(time.Now())
		milliseconds = t.millisecond
		seconds      = t.second
		minutes      = t.minute
		hours        = t.hour
		AM_or_PM     string
	)

	if hours < 12 {
		AM_or_PM = "AM"
	} else {
		AM_or_PM = "PM"
	}

	return utils.Stringify[any](hours, ":", minutes, ":", seconds, ".", milliseconds, " ", AM_or_PM)
}

func Current24Time() string {
	var (
		t            = getInstance(time.Now())
		milliseconds = t.millisecond
		seconds      = t.second
		minutes      = t.minute
		hours        = t.hour
	)

	return utils.Stringify[any](hours, ":", minutes, ":", seconds, ".", milliseconds)
}

// This will give the provided time in AM/PM
func GetATime(time_arg time.Time) string {
	var (
		t            = getInstance(time_arg)
		milliseconds = t.millisecond
		seconds      = t.second
		minutes      = t.minute
		hours        = t.hour
		AM_or_PM     string
	)

	if hours < 12 {
		AM_or_PM = "AM"
	} else {
		AM_or_PM = "PM"
	}

	return utils.Stringify[any](hours, ":", minutes, ":", seconds, ".", milliseconds, " ", AM_or_PM)
}

func Get24Time(time_arg time.Time) string {
	var (
		t            = getInstance(time_arg)
		milliseconds = t.millisecond
		seconds      = t.second
		minutes      = t.minute
		hours        = t.hour
	)

	return utils.Stringify[any](hours, ":", minutes, ":", seconds, ".", milliseconds)
}
