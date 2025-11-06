package api

import (
	"fmt"
	"io"
	"net/http"

	"cryptomaster.com/begin/datatypes"
)

var apiUrl = "https://jsonplaceholder.typicode.com/todos/1"

func GetRate(currency string) (*datatypes.Rate, error) {
	// upCurrency := strings.ToUpper(currency)
	res, err := http.Get(apiUrl)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json := string(bodyBytes)
		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("status code is : %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: 20}

	return &rate, nil
}
