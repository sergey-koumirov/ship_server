package messages

import(
	"binary"
)

type AddShip struct{
	login string
}

func (lr *AddShip) GetLogin() string {
	return lr.login
}

func GetAddShipFromBytes(bytes []byte) Message {
	m := new(AddShip)
	m.login = binary.ReadString(bytes)
	return m
}


