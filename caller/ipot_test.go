package caller

import (
	"context"
	"github.com/stretchr/testify/assert"
	"goshm/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	testCases := []struct {
		Name string
		NotMockServer bool
		MockRespBody []byte
		ExpectedResp *[]models.CallResponse
		ExpectedError bool
	}{
		{
			Name: "Success",
			MockRespBody: []byte("[[1595808000000,5075.0000,5150.0000,5025.0000,5125.0000,17645200]]"),
			ExpectedResp: &[]models.CallResponse{
				{
					Date:       1595808000000,
					OpenPrice:  "5075",
					ClosePrice: "5125",
					HighPrice:  "5150",
					LowPrice:   "5025",
				},
			},
		},
		{
			Name: "connection refused",
			NotMockServer: true,
			MockRespBody: []byte(""),
			ExpectedError: true,
		},
		{
			Name: "Unexpected body response",
			MockRespBody: []byte("unexpected response"),
			ExpectedError: true,
		},
		{
			Name: "unexpected idx 0",
			MockRespBody: []byte(`[["unexpected",5075.0000,5150.0000,5025.0000,5125.0000,17645200]]`),
			ExpectedError: true,
		},
		{
			Name: "unexpected idx 1",
			MockRespBody: []byte(`[[1595808000000,"unexpected",5150.0000,5025.0000,5125.0000,17645200]]`),
			ExpectedError: true,
		},
		{
			Name: "unexpected idx 2",
			MockRespBody: []byte(`[[1595808000000,5075.0000,"unexpected",5025.0000,5125.0000,17645200]]`),
			ExpectedError: true,
		},
		{
			Name: "unexpected idx 3",
			MockRespBody: []byte(`[[1595808000000,5075.0000,5150.0000,"unexpected",5125.0000,17645200]]`),
			ExpectedError: true,
		},
		{
			Name: "unexpected idx 4",
			MockRespBody: []byte(`[[1595808000000,5075.0000,5150.0000,5025.0000,"unexpected",17645200]]`),
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			baseUrl := "xxx"
			client := &http.Client{}
			if !tc.NotMockServer {
				callerMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write(tc.MockRespBody)
				}))
				baseUrl = callerMock.URL
				client = callerMock.Client()
			}

			timeNow := time.Now()
			ip := ipot{
				baseUrl: baseUrl,
			}
			resp, err := ip.Get(context.Background(), client, models.FetchArgs{
				Code:     "ASII",
				DateFrom: &timeNow,
				DateTo:   &timeNow,
			})

			if tc.ExpectedError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, *tc.ExpectedResp, resp)
		})
	}
}
