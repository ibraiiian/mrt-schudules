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
	StationId string 'json:"nid"'
	StationName string 'json:"title"'`
	ScheduleBundaraHI string 'json:"jadwal_h1biasa"'`
	ScheduleLebakBulus string 'json:"jadwal_lbubiasa"'`

}

type ScheduleResponse struct {
	StationId string 'json:"station"'
	Time string 'json:"time"'`