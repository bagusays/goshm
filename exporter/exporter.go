package exporter

import (
	"goshm/models"
)

const (
	OutputCSV  models.Output = "csv"
	OutputJSON models.Output = "json"
)

var Type = map[models.Output]Exporter{
	OutputCSV:  Csv(),
	OutputJSON: Json(),
}

type Exporter interface {
	Generate(param models.FetchArgs, data []models.CallResponse) error
	GetFilename(param models.FetchArgs) string
}
