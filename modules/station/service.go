package station

type station interface {
	GetStations() ([]StationResponse, error)
}

type service struct {
	client *http.Client
}

func NewService() service {
	return &service{
		client: &http.Client{},
		Timeout: 10* time.Second,
	}
}
 func (s *service) GetAllStations() ([]StationResponse, error) {
	url := "https://www.jakarta.go.id/metro/api/v1/stations"

	bytesResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []station
	err = json.Umarshal(bytesResponse, &stations)

	for _, item := range stations {
		response = append(response, StationResponse{
			ID: item.ID,
			Name: item.Name,
		})
	}
	)
	// kita keluarin response nya


	return

 }
