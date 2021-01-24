package models

import (
	"time"
)

type Output string

type FetchArgs struct {
	Code     string
	DateFrom *time.Time
	DateTo   *time.Time
}

type CallResponse struct {
	Date       int64
	OpenPrice  string
	ClosePrice string
	HighPrice  string
	LowPrice   string
}
