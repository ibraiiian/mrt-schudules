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

	// Mock data
	response = append(response, StationResponse{
		ID:   "1",
		Name: "Kota",
	})
	response = append(response, StationResponse{
		ID:   "2",
		Name: "Bundaran HI",
	})
	response = append(response, StationResponse{
		ID:   "3",
		Name: "Senayan",
	})
	response = append(response, StationResponse{
		ID:   "4",
		Name: "Istora Mandiri",
	})
	response = append(response, StationResponse{
		ID:   "5",
		Name: "Blok M",
	})

	return response, nil
}

func (s *Service) GetScheduleByStationID(id string) ([]ScheduleResponse, error) {
	var response []ScheduleResponse

	// Mock schedule data berdasarkan station ID
	schedules := map[string][]string{
		"1": {"06:00", "06:15", "06:30", "06:45", "07:00", "07:15", "07:30", "08:00"},
		"2": {"06:05", "06:20", "06:35", "06:50", "07:05", "07:20", "07:35", "08:05"},
		"3": {"06:10", "06:25", "06:40", "06:55", "07:10", "07:25", "07:40", "08:10"},
		"4": {"06:12", "06:27", "06:42", "06:57", "07:12", "07:27", "07:42", "08:12"},
		"5": {"06:18", "06:33", "06:48", "07:03", "07:18", "07:33", "07:48", "08:18"},
	}

	times, ok := schedules[id]
	if !ok {
		times = schedules["1"] // default ke station 1
	}

	for _, time := range times {
		response = append(response, ScheduleResponse{
			StationID: id,
			Time:      time,
		})
	}

	return response, nil
}
