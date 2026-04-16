package station

type Station struct {
	ID   string `json:"id"`
	Name string `json:"title"`
}

type StationResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	StationID     string `json:"station_id"`
	StationName   string `json:"station_name"`
	ScheduleLebak string `json:"schedule_lebak"`
	ScheduleBund  string `json:"schedule_bundara"`
}

type ScheduleResponse struct {
	StationID string `json:"station_id"`
	Time      string `json:"time"`
}
