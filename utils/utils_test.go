package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFromArgToUnixDateFormatter(t *testing.T) {
	reqDate := "30-12-2020"
	epoch, err := FromArgToUnixDateFormatter(reqDate)
	assert.NoError(t, err)
	assert.Equal(t, int64(1609286400), epoch.Unix())
}

func TestToCallerDateFormatter(t *testing.T) {
	reqDate := "30-12-2020"
	dateInTime, err := time.Parse("02-01-2006", reqDate)
	assert.NoError(t, err)
	str := ToCallerDateFormatter(&dateInTime)
	assert.Equal(t, "12/30/2020", str)
}

func TestToReadableDate(t *testing.T) {
	reqDate := "30/12/2020"
	dateInTime, err := time.Parse("02/01/2006", reqDate)
	assert.NoError(t, err)
	str := ToReadableDate(&dateInTime)
	assert.Equal(t, "30-12-2020", str)
}

func TestToWriterDate(t *testing.T) {
	reqDate := "30-12-2020"
	dateInTime, err := time.Parse("02-01-2006", reqDate)
	assert.NoError(t, err)
	str := ToWriterDate(&dateInTime)
	assert.Equal(t, "30122020", str)
}

func TestUnixToDate(t *testing.T) {
	dateInTime := UnixToDate(int64(1609286400))
	reqDate := "30-12-2020"
	expectedDateInTime, err := time.Parse("02-01-2006", reqDate)
	assert.NoError(t, err)
	assert.Equal(t, expectedDateInTime.Unix(), dateInTime.Unix())
}
