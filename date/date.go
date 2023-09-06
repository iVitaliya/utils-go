package date

import (
	"time"

	"github.com/iVitaliya/utils-go/utils"
)

type dateInstance struct {
	ToFormatted string
	Object      struct {
		Time  string
		Day   dayInstance
		Month monthInstance
		Year  int
	}
}

func CurrentDate() *dateInstance {
	var (
		t     = CurrentATime()
		day   = CurrentDay()
		month = CurrentMonth()
		year  = time.Now().Local().Year()

		formatted = utils.Stringify[any]("the ", day.Human, " of ", month.Human, ", ", year, " @ ", t)
	)

	return &dateInstance{
		ToFormatted: formatted,
		Object: struct {
			Time  string
			Day   dayInstance
			Month monthInstance
			Year  int
		}{
			Time:  t,
			Day:   *day,
			Month: *month,
			Year:  year,
		},
	}
}

func GetDate(time_arg time.Time, am_or_pm_format bool, day int, month int, year int) *dateInstance {
	var (
		t string
		d = GetDay(day)
		m = GetMonth(month)
		y int

		formatted string
	)

	if year < 1970 {
		y = 1970
	} else if year > time.Now().Local().Year() {
		y = time.Now().Local().Year()
	} else {
		y = year
	}

	if am_or_pm_format {
		formatted = utils.Stringify[any]("the ", d.Human, " of ", m.Human, ", ", y, " @ ", GetATime(time_arg))
	} else {
		formatted = utils.Stringify[any]("the ", d.Human, " of ", m.Human, ", ", y, " @ ", Get24Time(time_arg))
	}

	return &dateInstance{
		ToFormatted: formatted,
		Object: struct {
			Time  string
			Day   dayInstance
			Month monthInstance
			Year  int
		}{
			Time:  t,
			Day:   *d,
			Month: *m,
			Year:  y,
		},
	}
}
