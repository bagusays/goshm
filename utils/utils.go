package utils

import (
	"errors"
	"time"
)

var argsExpectedFormatDate = "02-01-2006"
var callerFormatDate = "01/02/2006"

func FromArgToUnixDateFormatter(date string) (*time.Time, error) {
	dateInTime, err := time.Parse(argsExpectedFormatDate, date)
	if err != nil {
		return nil, errors.New("invalid date format, please use dd-mm-yyyy")
	}
	return &dateInTime, nil
}

func ToCallerDateFormatter(date *time.Time) string {
	return date.Format(callerFormatDate)
}

func ToWriterDate(date *time.Time) string {
	writerFormatDate := "02012006"
	return date.Format(writerFormatDate)
}

func UnixToDate(date int64) *time.Time {
	unix := time.Unix(date, 0)
	return &unix
}

func ToReadableDate(date *time.Time) string {
	return date.Format(argsExpectedFormatDate)
}
