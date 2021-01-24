package writer

import (
	"encoding/csv"
	"fmt"
	"goshm/models"
	"goshm/utils"
	"os"
)

type csvWriter struct {
	filename string
}

func Csv() Writer {
	return &csvWriter{}
}

func (c *csvWriter) Generate(param models.FetchArgs, data []models.CallResponse) error {
	fmt.Printf("Writing %s into csv... ", param.Code)

	records := [][]string{
		{"date", "open", "close", "low", "high"},
	}
	for _, d := range data {
		date := utils.ToReadableDate(utils.UnixToDate(d.Date/1000))
		tmp := []string{date, d.OpenPrice, d.ClosePrice, d.LowPrice, d.HighPrice}
		records = append(records, tmp)
	}

	f, err := os.Create(c.GetFilename(param))
	defer f.Close()
	if err != nil {
		fmt.Print("[FAILED]\n")
		return err
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)
	if err != nil {
		fmt.Print("[FAILED]\n")
		return err
	}

	fmt.Print("[OK]\n")
	return nil
}

func (c *csvWriter) GetFilename(param models.FetchArgs) string {
	if c.filename != "" {
		return c.filename
	}
	return fmt.Sprintf("%s#%s-%s.csv", param.Code, utils.ToWriterDate(param.DateFrom), utils.ToWriterDate(param.DateTo))
}