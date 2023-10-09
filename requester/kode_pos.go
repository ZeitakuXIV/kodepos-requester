package requester

import (
	"encoding/json"
	"go-api/model"
	"net/http"
)

type RequestKodePos struct {
	client *http.Client
}

func NewRequestKodePos(client *http.Client) RequestKodePos {
	return RequestKodePos{client: client}
}

func (R RequestKodePos) FetchKodePos() ([]model.KodePos, error) {
	var data []model.KodePos
	request, err := http.NewRequest("GET", model.BaseUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := R.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
