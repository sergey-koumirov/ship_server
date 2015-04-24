package abyss

import(
	"time"
	"binary"
	"messages"
	//"fmt"
	"fmt"
)

type WorldImage struct {
	TickTime time.Time
	objects []*Object
	dumper binary.Dumper
}

func (wi *WorldImage) ToBytes() []byte {
	wi.dumper.Reset()
	wi.dumper.PutTSTime(wi.TickTime)

	wi.dumper.PutInt8(binary.OBJECTS_TABLE_CODE)
	wi.dumper.PutInt16( int16( len(wi.objects) ) )
	for _, v := range wi.objects {
		v.ToBytes(&wi.dumper)
	}

	return wi.dumper.Bytes()
}

func (wi *WorldImage) MoveObjects(delta float64){
	for _, v := range wi.objects {
		 v.ChangeState(delta)
	}
}

func (wi *WorldImage) AddShip(owner string) {
	exists := false
	for _, v := range wi.objects {
		if v.owner == owner {
			exists = true
		}
	}
	if !exists {
		wi.objects = append(wi.objects, &Object{typeID: SHIP_TYPE_ID, x: 100, y: 100, angle: 0, speedX: 0, speedY: 0, acceleration:0, rotation: 0, owner: owner} )
	}
}

func (wi *WorldImage) UpdateShipControl(sd *messages.ShipDriving) {
	for _, v := range wi.objects {
		if v.owner == sd.GetLogin(){
			v.rotation = sd.GetRotation()
			v.acceleration = sd.GetAcceleration()
		}
	}
}

func (wi *WorldImage) LoadShipTemplates(){
	var paths []string = []string{
	"../ShipJ/res/Dummy.json",
	"../ShipJ/res/Box.json",
	}

	for _,v := range paths{
		st := LoadShipTemplate(v)
		fmt.Println(st)
	}
}
