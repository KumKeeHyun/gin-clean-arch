package model

type Loc struct {
	Lat float64
	Lon float64
}

type Node struct {
	Name     string
	Group    string
	Location Loc
}

type Sensor struct {
	Name       string
	ValueNames []string
}
