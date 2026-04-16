package station

import (
	"net/http"
)

type Service struct {
	client *http.Client
}

func NewService() *Service {
	return &Service{
		client: &http.Client{},
	}
}

func (s *Service) GetAllStations() ([]StationResponse, error) {
	var response []StationResponse

	// Mock data for now
	response = append(response, StationResponse{
		ID:   "1",
		Name: "Kota",
	})
	response = append(response, StationResponse{
		ID:   "2",
		Name: "Bundaran HI",
	})

	return response, nil
}
