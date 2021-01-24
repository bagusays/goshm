package executors

import (
	"errors"
	"goshm/mocks"
	"goshm/models"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGoShm_Fetch(t *testing.T) {
	timeNow := time.Now()
	testCases := []struct {
		Name           string
		MockCallerResp []models.CallResponse
		MockCallerErr  error
		MockWriterErr  error
		ParamFetchArgs models.FetchArgs
		ExpectedError  bool
	}{
		{
			Name: "Success",
			MockCallerResp: []models.CallResponse{
				{
					Date:       1595808000000,
					OpenPrice:  "5075",
					ClosePrice: "5125",
					HighPrice:  "5150",
					LowPrice:   "5025",
				},
			},
			MockCallerErr: nil,
			MockWriterErr: nil,
			ParamFetchArgs: models.FetchArgs{
				Code:     "ASII",
				DateFrom: &timeNow,
				DateTo:   &timeNow,
			},
		},
		{
			Name:           "Caller error",
			MockCallerResp: nil,
			MockCallerErr:  errors.New("unexpected error"),
			MockWriterErr:  nil,
			ParamFetchArgs: models.FetchArgs{
				Code:     "ASII",
				DateFrom: &timeNow,
				DateTo:   &timeNow,
			},
			ExpectedError: true,
		},
		{
			Name: "Writer error",
			MockCallerResp: []models.CallResponse{
				{
					Date:       1595808000000,
					OpenPrice:  "5075",
					ClosePrice: "5125",
					HighPrice:  "5150",
					LowPrice:   "5025",
				},
			},
			MockCallerErr: nil,
			MockWriterErr: errors.New("unexpected error"),
			ParamFetchArgs: models.FetchArgs{
				Code:     "ASII",
				DateFrom: &timeNow,
				DateTo:   &timeNow,
			},
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockFetch := &mocks.Caller{}
			mockFetch.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(tc.MockCallerResp, tc.MockCallerErr)

			mockWriter := &mocks.Writer{}
			mockWriter.On("Generate", mock.Anything, mock.Anything).Return(tc.MockWriterErr)

			shm := New()
			err := shm.Fetch(mockFetch, tc.ParamFetchArgs, mockWriter)

			if tc.ExpectedError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestGoShm_SetHttpClient(t *testing.T) {
	expectedHttpClient := &http.Client{Timeout: 99}
	shm := New()
	shm.SetHttpClient(expectedHttpClient)
	assert.Equal(t, expectedHttpClient, shm.httpClient)
}
