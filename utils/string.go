package utils

import "fmt"

func Stringify[T any](value ...T) string {
	var s string

	for _, val := range value {
		s = s + fmt.Sprint(val)
	}

	return s
}

func ForEachString(arr []string) {}
