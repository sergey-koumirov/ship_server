package abyss

import (
	"math"
	"binary"
//	"fmt"
)

const SHIP_TYPE_ID = int16(1)

type Object struct{
	typeID int16
	x,y,speedX,speedY,angle,acceleration,rotation float64
	owner string
}

func (o *Object) ChangeState(delta float64){
	o.angle = o.angle + o.rotation / 1000.0 * delta

	o.speedX = o.speedX + o.acceleration * math.Cos( math.Pi * o.angle / 180 ) / 1000.0 * delta
	o.speedY = o.speedY + o.acceleration * math.Sin( math.Pi * o.angle / 180 ) / 1000.0 * delta

	o.x = o.x + (o.speedX * delta)
	o.y = o.y + (o.speedY * delta)
}

func (o *Object) ToBytes(dumper *binary.Dumper){
	dumper.PutInt16(o.typeID)
	dumper.PutFloat64(o.x)
	dumper.PutFloat64(o.y)
	dumper.PutFloat64(o.speedX)
	dumper.PutFloat64(o.speedY)
	dumper.PutFloat64(o.angle)
	dumper.PutFloat64(o.acceleration)
	dumper.PutFloat64(o.rotation)
	dumper.PutString(o.owner)
}
