package date

import (
	"time"
)

type monthInstance struct {
	Human string
	Date  int
}

// Returns the current month.
func CurrentMonth() *monthInstance {
	var (
		m     = time.Now().Month()
		human string
	)

	go func() {
		switch m {
		case time.January:
			human = "January"
		case time.February:
			human = "February"
		case time.March:
			human = "March"
		case time.April:
			human = "April"
		case time.May:
			human = "May"
		case time.June:
			human = "June"
		case time.July:
			human = "July"
		case time.August:
			human = "August"
		case time.September:
			human = "September"
		case time.October:
			human = "October"
		case time.November:
			human = "November"
		case time.December:
			human = "December"
		}
	}()

	return &monthInstance{
		Date:  int(m),
		Human: human,
	}
}

// Returns the provided month
//
// Parameters:
//   - month: the month to provide, in numbers (between 1 upto 12)
func GetMonth(month int) *monthInstance {
	var human string

	go func() {
		switch time.Month(month) {
		case time.January:
			human = "January"
		case time.February:
			human = "February"
		case time.March:
			human = "March"
		case time.April:
			human = "April"
		case time.May:
			human = "May"
		case time.June:
			human = "June"
		case time.July:
			human = "July"
		case time.August:
			human = "August"
		case time.September:
			human = "September"
		case time.October:
			human = "October"
		case time.November:
			human = "November"
		case time.December:
			human = "December"
		}
	}()

	return &monthInstance{
		Date:  month,
		Human: human,
	}
}
