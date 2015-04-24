package binary

import(
	b "encoding/binary"
)

func ReadString(bytes []byte) string {
	length := b.LittleEndian.Uint16(bytes)
	return string( bytes[2:length+2] )
}

func ReadUInt16(bytes []byte) uint16 {
	return b.LittleEndian.Uint16(bytes)
}

func ReadInt16(bytes []byte) int16 {
	return int16(bytes[0]) + int16(bytes[1]) <<8
}
