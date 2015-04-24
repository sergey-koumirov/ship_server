package abyss

import (
	"io/ioutil"
	"encoding/json"
)

type Point struct{
	X,Y float64
}

type Plate struct{
	Material string
	Points []Point
}

type ShipTemplate struct{
	TypeID int64
	Plates []Plate
}

func LoadShipTemplate(path string) ShipTemplate {
	bytes, _ := ioutil.ReadFile(path)
	var ship ShipTemplate
	_ = json.Unmarshal(bytes, &ship)
	return ship
}
