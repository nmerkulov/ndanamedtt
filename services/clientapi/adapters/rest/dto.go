package rest

import "ndanamedtt/services/clientapi/domain"

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type Next int

type PortListResponse struct {
	Items []Port
	//id-based pagination. useful to avoid "limit offset" pagination
	Next Next
}

func portToRest(p domain.Port) Port {
	return Port(p)
}
