package messages

import(
	"binary"
)

type LoginRequest struct{
	login string
}

func (lr *LoginRequest) GetLogin() string {
	return lr.login
}

func GetLoginRequestFromBytes(bytes []byte) Message {
	m := new(LoginRequest)
	m.login = binary.ReadString(bytes)
	return m
}


