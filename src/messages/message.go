package messages

const LOGIN_REQUEST_CODE = uint16(1)
const SHIP_DRIVING_CODE =  uint16(2)
const ADD_SHIP_CODE =      uint16(3)


type Message interface{
    GetLogin() string
}

