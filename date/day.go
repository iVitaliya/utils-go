package date

import (
	"fmt"
	"time"
)

type dayInstance struct {
	Human string
	Date  int
}

// Returns the current day
func CurrentDay() *dayInstance {
	var (
		d     = time.Now().Day()
		human string
	)

	go func() {
		switch d {
		case 1:
		case 21:
		case 31:
			human = fmt.Sprint(d) + "st"
		case 2:
		case 22:
			human = fmt.Sprint(d) + "nd"
		case 3:
		case 23:
			human = fmt.Sprint(d) + "rd"
		default:
			human = fmt.Sprint(d) + "th"
		}
	}()

	return &dayInstance{
		Date:  d,
		Human: human,
	}
}

// Returns the provided day
//
// Parameters:
//   - day: the day to provide, in numbers (between 1 upto 31)
func GetDay(day int) *dayInstance {
	var human string

	go func() {
		switch day {
		case 1:
		case 21:
		case 31:
			human = fmt.Sprint(day) + "st"
		case 2:
		case 22:
			human = fmt.Sprint(day) + "nd"
		case 3:
		case 23:
			human = fmt.Sprint(day) + "rd"
		default:
			human = fmt.Sprint(day) + "th"
		}
	}()

	return &dayInstance{
		Date:  day,
		Human: human,
	}
}
