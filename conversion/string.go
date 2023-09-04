package conversion

import "fmt"

func ToString[T any](value T) string {
	return fmt.Sprint(value)
}

func DayStr(day int, suffix string) string {
	
}