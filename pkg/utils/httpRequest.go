package utils

import (
	"encoding/json"
	"fmt"
	"go-medium-shapes/pkg/constants"
	"io"
	"net/http"
)

type ResponseGet struct {
	Data struct {
		Email string `json:"email"`
	}
}

func HttpGetUserData(id string) (ResponseGet, error) {
	var userData ResponseGet

	url := constants.URL_API_USERS + id
	httpResponse, err := http.Get(url)
	if err != nil {
		return userData, fmt.Errorf("HttpGetUserData. Error consultando datos de api (%s). %s", url, err)
	}

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return userData, fmt.Errorf("HttpGetUserData. Error leyendo datos de response. %s", err)
	}

	err = json.Unmarshal(body, &userData)
	if err != nil {
		return userData, fmt.Errorf("HttpGetUserData. Error parseando datos de response. %s", err)
	}

	return userData, nil
}
