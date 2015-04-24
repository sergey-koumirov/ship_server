package messages

import(
	"binary"
)

type ShipDriving struct{
	login string
	acceleration int16
	rotation int16
}

func (lr *ShipDriving) GetLogin() string {
	return lr.login
}

func (lr *ShipDriving) GetAcceleration() float64 {
	return float64(lr.acceleration)
}

func (lr *ShipDriving) GetRotation() float64 {
	return float64(lr.rotation)
}

func GetShipDrivingFromBytes(bytes []byte) Message {
	var shift int16 = 0

	m := new(ShipDriving)

	m.login = binary.ReadString(bytes[shift:])
	shift = shift + 2 + int16( len(m.login) )

	m.acceleration = binary.ReadInt16(bytes[shift:])
	shift = shift + 2

	m.rotation = binary.ReadInt16(bytes[shift:])

	return m
}


