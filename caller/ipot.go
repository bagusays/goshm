package caller

import (
	"context"
	"encoding/json"
	"fmt"
	"goshm/models"
	"goshm/utils"
	"net/http"
	"strconv"
)

type ipot struct {
	baseUrl string
}

func Ipot() Caller {
	return &ipot{
		baseUrl: "https://www.indopremier.com",
	}
}

func (i *ipot) Get(ctx context.Context, httpClient *http.Client, args models.FetchArgs) ([]models.CallResponse, error) {
	fmt.Printf("Fetching %s from database... ", args.Code)
	dateFrom := utils.ToCallerDateFormatter(args.DateFrom)
	dateTo := utils.ToCallerDateFormatter(args.DateTo)

	baseUrl := fmt.Sprintf(`%s/module/saham/include/json-charting.php?code=%s&start=%s&end=%s`, i.baseUrl, args.Code, dateFrom, dateTo)

	req, err := http.NewRequestWithContext(ctx, "GET", baseUrl, nil)
	if err != nil {
		fmt.Print("[FAILED]\n")
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Print("[FAILED]\n")
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	var respBody [][]json.RawMessage
	err = json.NewDecoder(res.Body).Decode(&respBody)
	if err != nil {
		fmt.Print("[FAILED]\n")
		return nil, err
	}

	var result []models.CallResponse
	for _, data := range respBody {
		time, err := strconv.ParseInt(string(data[0]), 10, 64)
		if err != nil {
			return nil, err
		}
		openPrice, err := strconv.ParseFloat(string(data[1]), 64)
		if err != nil {
			return nil, err
		}
		closePrice, err := strconv.ParseFloat(string(data[4]), 64)
		if err != nil {
			return nil, err
		}
		lowPrice, err := strconv.ParseFloat(string(data[3]), 64)
		if err != nil {
			return nil, err
		}
		highPrice, err := strconv.ParseFloat(string(data[2]), 64)
		if err != nil {
			return nil, err
		}
		result = append(result, models.CallResponse{
			Date:   time,
			OpenPrice:  strconv.FormatFloat(openPrice, 'f', 0, 64),
			ClosePrice: strconv.FormatFloat(closePrice, 'f', 0, 64),
			HighPrice:  strconv.FormatFloat(highPrice, 'f', 0, 64),
			LowPrice:   strconv.FormatFloat(lowPrice, 'f', 0, 64),
		})
	}

	fmt.Print("[OK]\n")

	return result, nil
}
