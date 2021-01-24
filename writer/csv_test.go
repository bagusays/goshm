package writer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goshm/models"
	"os"
	"testing"
	"time"
)

func TestGenerateCSV(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		dateFrom, err := getWriterDate("30-12-2020")
		assert.NoError(t, err)
		dateTo, err := getWriterDate("30-12-2020")
		assert.NoError(t, err)

		param := models.FetchArgs{
			Code:     "ASII",
			DateFrom: dateFrom,
			DateTo:   dateTo,
		}
		generator := Csv()
		err = generator.Generate(param, []models.CallResponse{})
		assert.NoError(t, err)

		filename := fmt.Sprintf("%s#%s-%s.csv", param.Code, "30122020", "30122020")
		_, err = os.Stat(filename)
		assert.Equal(t, false, os.IsNotExist(err))
		err = os.Remove(generator.GetFilename(param))
		assert.NoError(t, err)
	})

	t.Run("Failed write file", func(t *testing.T) {
		dateFrom, err := getWriterDate("30-12-2020")
		assert.NoError(t, err)
		dateTo, err := getWriterDate("30-12-2020")
		assert.NoError(t, err)

		param := models.FetchArgs{
			Code:     "ASII",
			DateFrom: dateFrom,
			DateTo:   dateTo,
		}
		generator := csvWriter{filename: "/"}
		err = generator.Generate(param, []models.CallResponse{})
		assert.Error(t, err)
	})

	t.Run("Failed write records", func(t *testing.T) {
		dateFrom, err := getWriterDate("30-12-2020")
		assert.NoError(t, err)
		dateTo, err := getWriterDate("30-12-2020")
		assert.NoError(t, err)

		param := models.FetchArgs{
			Code:     "ASII",
			DateFrom: dateFrom,
			DateTo:   dateTo,
		}
		generator := csvWriter{filename: "/"}
		err = generator.Generate(param, []models.CallResponse{})
		assert.Error(t, err)
	})
}

func getWriterDate(date string) (*time.Time, error) {
	reqDate := date
	dateInTime, err := time.Parse("02-01-2006", reqDate)
	if err != nil {
		return nil, err
	}
	return &dateInTime, nil
}