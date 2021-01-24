package exporter

import (
	"fmt"
	"goshm/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonWriter_Generate(t *testing.T) {
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

		generator := Json()
		err = generator.Generate(param, []models.CallResponse{})
		assert.NoError(t, err)

		filename := fmt.Sprintf("%s#%s-%s.json", param.Code, "30122020", "30122020")
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
}
