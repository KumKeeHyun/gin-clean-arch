package presenter

import "github.com/KumKeeHyun/gin-clean-arch/app/domain/model"

type Node struct {
	UUID     string         `json:"uuid"`
	Name     string         `json:"name"`
	Location string         `json:"location"`
	Sensors  []model.Sensor `json:"sensors"`
}

func ToNodes(n []model.Node) []Node {
	res := make([]Node, len(n))
	for i, node := range n {
		res[i] = ToNode(&node)
	}
	return res
}

func ToNode(n *model.Node) Node {
	return Node{
		UUID:     n.UUID,
		Name:     n.Name,
		Location: n.Location,
		Sensors:  make([]model.Sensor, 0),
	}
}

// type Sensor struct {
// 	UUID      string        `json:"uuid"`
// 	Name      string        `json:"name"`
// 	ValueList []SensorValue `json:"value_list"`
// }

// func ToSensors(s []model.Sensor) []Sensor {
// 	res := make([]Sensor, len(s))
// 	for i, sensor := range s {
// 		res[i] = ToSensor(&sensor)
// 	}
// 	return res
// }

// func ToSensor(s *model.Sensor) Sensor {
// 	return Sensor{
// 		UUID:      s.UUID,
// 		Name:      s.Name,
// 		ValueList: make([]SensorValue, 0),
// 	}
// }

// type SensorValue struct {
// 	ValueName string `json:"value_name"`
// }

// func ToSensorValue(sv []model.SensorValue) []SensorValue {
// 	res := make([]SensorValue, len(sv))
// 	for i, sn := range sv {
// 		res[i] = SensorValue{
// 			ValueName: sn.ValueName,
// 		}
// 	}
// 	return res
// }
