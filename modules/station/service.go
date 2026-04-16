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
	// hit url

	// kita keluarin response nya

	
	return

 }
