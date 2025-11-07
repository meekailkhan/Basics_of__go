package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cryptomaster.com/begin/datatypes"
)

var apiUrl = "https://jsonplaceholder.typicode.com/todos/%v"

func GetRate(userId string) (*datatypes.UserResponse, error) {
	// upCurrency := strings.ToUpper(currency)
	fetchUrl := fmt.Sprintf(apiUrl, userId)
	res, err := http.Get(fetchUrl)

	if err != nil {
		return nil, err
	}
	var respose datatypes.UserResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &respose)
		if err != nil {
			return nil, err
		}
		// fmt.Println(respose)

	} else {
		return nil, fmt.Errorf("status code is : %v", res.StatusCode)
	}

	// rate := datatypes.Rate{}

	return &respose, nil
}
