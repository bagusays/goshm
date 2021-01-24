package caller

import (
	"context"
	"goshm/models"
	"net/http"
)

type Caller interface {
	Get(ctx context.Context, httpClient *http.Client, args models.FetchArgs) ([]models.CallResponse, error)
}
