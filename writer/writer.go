package writer

import (
	"goshm/models"
)

const (
	OutputCSV models.Output = "OutputCSV"
)

var Generator = map[models.Output]Writer{
	OutputCSV: Csv(),
}

type Writer interface {
	Generate(param models.FetchArgs, data []models.CallResponse) error
	GetFilename(param models.FetchArgs) string
}
