package executors

import (
	"context"
	"goshm/caller"
	"goshm/exporter"
	"goshm/models"
	"net/http"
	"time"
)

type GoShm struct {
	httpClient *http.Client
}

func New() *GoShm {
	return &GoShm{httpClient: &http.Client{Timeout: 5 * time.Second}}
}

func (n *GoShm) Fetch(platform caller.Caller, args models.FetchArgs, writer exporter.Exporter) error {
	resp, err := platform.Get(context.Background(), n.httpClient, args)
	if err != nil {
		return err
	}

	err = writer.Generate(args, resp)
	if err != nil {
		return err
	}
	return nil
}

func (n *GoShm) SetHttpClient(client *http.Client) *GoShm {
	n.httpClient = client
	return n
}
