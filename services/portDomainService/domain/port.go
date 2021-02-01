package domain

type PortID int

type Port struct {
	ID          PortID
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}

func (pID PortID) IsSet() bool {
	return pID == 0
}
