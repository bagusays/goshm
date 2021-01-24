package exporter

import (
	"encoding/json"
	"fmt"
	"goshm/models"
	"goshm/utils"
	"io/ioutil"
)

type jsonExporter struct {
	Date       string `json:"date"`
	OpenPrice  string `json:"open"`
	ClosePrice string `json:"close"`
	HighPrice  string `json:"high"`
	LowPrice   string `json:"low"`
}

type jsonWriter struct {
	filename string
}

func Json() Exporter {
	return &jsonWriter{}
}

func (j *jsonWriter) Generate(param models.FetchArgs, data []models.CallResponse) error {
	fmt.Printf("Writing %s into json... ", param.Code)

	var records []jsonExporter
	for _, d := range data {
		date := utils.ToReadableDate(utils.UnixToDate(d.Date / 1000))
		records = append(records, jsonExporter{
			Date:       date,
			OpenPrice:  d.OpenPrice,
			ClosePrice: d.ClosePrice,
			HighPrice:  d.HighPrice,
			LowPrice:   d.LowPrice,
		})
	}
	jsonRecords, err := json.Marshal(records)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(j.GetFilename(param), jsonRecords, 0644)
	if err != nil {
		fmt.Print("[FAILED]\n")
		return err
	}

	fmt.Print("[OK]\n")
	return nil
}

func (j *jsonWriter) GetFilename(param models.FetchArgs) string {
	if j.filename != "" {
		return j.filename
	}
	return fmt.Sprintf("%s#%s-%s.json", param.Code, utils.ToWriterDate(param.DateFrom), utils.ToWriterDate(param.DateTo))
}
