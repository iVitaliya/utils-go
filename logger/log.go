package logger

import (
	"fmt"
	"strings"

	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/utils-go/date"
)

type loggerInstance struct{}

func NewLogger() *loggerInstance {
	return &loggerInstance{}
}

func (log *loggerInstance) Info(text ...string) {
	var (
		dt      = date.CurrentDate().ToFormatted
		d       = colors.BrightBlack("[") + colors.Dim(colors.BrightYellow(dt)) + colors.BrightBlack("]")
		stt     = colors.Gray("(") + colors.BrightBlue("INFO") + colors.Gray(")")
		lastInd = len(text) - 1
		str     string
	)

	for i := 0; i < len(text); i++ {
		if i == lastInd {
			str = str + text[i]
		}

		str = str + text[i] + " "
	}

	fmt.Print(strings.Join([]string{d, stt, ":", str}, " "))
}

func (log *loggerInstance) Debug(text ...string) {
	var (
		dt      = date.CurrentDate().ToFormatted
		d       = colors.BrightBlack("[") + colors.Dim(colors.BrightYellow(dt)) + colors.BrightBlack("]")
		stt     = colors.Gray("(") + colors.Green("DEBUG") + colors.Gray(")")
		lastInd = len(text) - 1
		str     string
	)

	for i := 0; i < len(text); i++ {
		if i == lastInd {
			str = str + text[i]
		}

		str = str + text[i] + " "
	}

	fmt.Print(strings.Join([]string{d, stt, ":", str}, " "))
}

func (log *loggerInstance) Warn(text ...string) {
	var (
		dt      = date.CurrentDate().ToFormatted
		d       = colors.BrightBlack("[") + colors.Dim(colors.BrightYellow(dt)) + colors.BrightBlack("]")
		stt     = colors.Gray("(") + colors.Dim(colors.BrightYellow("WARN")) + colors.Gray(")")
		lastInd = len(text) - 1
		str     string
	)

	for i := 0; i < len(text); i++ {
		if i == lastInd {
			str = str + text[i]
		}

		str = str + text[i] + " "
	}

	fmt.Print(strings.Join([]string{d, stt, ":", str}, " "))
}

func (log *loggerInstance) Error(text ...string) {
	var (
		dt      = date.CurrentDate().ToFormatted
		d       = colors.BrightBlack("[") + colors.Dim(colors.BrightYellow(dt)) + colors.BrightBlack("]")
		stt     = colors.Gray("(") + colors.Red("ERROR") + colors.Gray(")")
		lastInd = len(text) - 1
		str     string
	)

	for i := 0; i < len(text); i++ {
		if i == lastInd {
			str = str + text[i]
		}

		str = str + text[i] + " "
	}

	fmt.Print(strings.Join([]string{d, stt, ":", str}, " "))
}
