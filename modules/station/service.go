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
	//layer service

	return

 }
