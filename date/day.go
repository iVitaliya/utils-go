package date

import "time"

type day struct {
	Human string
	Date int
}

// Returns the current day.
func CurrentDay() *day {
	var (
		d = time.Now().Day()
		human string
	)

	go func(
		switch d {
		case 1:
		case 21:
		case 31:
			human = 
		}
	)()

	return &day{
		Date: d,
	}
}