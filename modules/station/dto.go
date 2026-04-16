package station

type Station struct { 
	id string 'json:id:"nid"'
	name string 'json:"title"'
}

struct StationResponse {
	id string 'json:id:"id"'
	name string 'json:"name"'