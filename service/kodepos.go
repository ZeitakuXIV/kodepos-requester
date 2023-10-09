package service

import (
	"go-api/model"
	"go-api/requester"
	"net/http"
)

type KodePosService struct {
	client *http.Client
}

func NewKodePosService(client *http.Client) KodePosService {
	return KodePosService{
		client: client,
	}
}

func (ks KodePosService) GetAllKodePos() ([]model.KodePos, error) {
	kodePosRequester := requester.NewRequestKodePos(ks.client)
	kodePos, err := kodePosRequester.FetchKodePos()
	if err != nil {
		return nil, err
	}
	return kodePos, nil
}
