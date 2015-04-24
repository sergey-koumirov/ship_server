package binary

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"time"
)

const SERVER_TIME_CODE = int8(1)
const OBJECTS_TABLE_CODE = int8(2)

type Dumper struct{
	buffer bytes.Buffer
}

func (d *Dumper) Bytes() []byte {
	return d.buffer.Bytes()
}

func (d *Dumper) Reset() {
	d.buffer.Reset()
}

func (d *Dumper) PutInt8(v int8){
	err := binary.Write(&d.buffer, binary.LittleEndian, v)
	if err != nil { fmt.Println("binary.Write failed:", err)}
}

func (d *Dumper) PutInt16(v int16){
	err := binary.Write(&d.buffer, binary.LittleEndian, v)
	if err != nil { fmt.Println("binary.Write failed:", err)}
}

func (d *Dumper) PutInt64(v int64){
	err := binary.Write(&d.buffer, binary.LittleEndian, v)
	if err != nil { fmt.Println("binary.Write failed:", err)}
}

func (d *Dumper) PutFloat64(v float64){
	err := binary.Write(&d.buffer, binary.LittleEndian, v)
	if err != nil { fmt.Println("binary.Write failed:", err)}
}

func (d *Dumper) PutTime(v time.Time){
	err := binary.Write(&d.buffer, binary.LittleEndian, v.Unix())
	if err != nil { fmt.Println("binary.Write failed:", err)}
}

func (d *Dumper) PutString(v string){
	d.PutInt16( int16( len(v) ) )
	err := binary.Write(&d.buffer, binary.LittleEndian, []byte(v))
	if err != nil { fmt.Println("binary.Write failed:", err)}
}

func (d *Dumper) PutTSTime(v time.Time){
	err := binary.Write(&d.buffer, binary.LittleEndian, SERVER_TIME_CODE)
	if err != nil { fmt.Println("binary.Write failed:", err)}

	d.PutTime(v)
}





//buf := new(bytes.Buffer)
//
//var d1 int64 = 65
//var d2 float64 = 65
//var d3 time.Time = time.Now()
//var d4 string = "65"
//
//
//err := binary.Write(buf, binary.LittleEndian, d1)
//if err != nil { fmt.Println("binary.Write failed:", err)}
//
//err = binary.Write(buf, binary.LittleEndian, d2)
//if err != nil { fmt.Println("binary.Write failed:", err)}
//
//err = binary.Write(buf, binary.LittleEndian, d3.Unix())
//if err != nil { fmt.Println("binary.Write failed:", err)}
//
//err = binary.Write(buf, binary.LittleEndian, []byte(d4))
//if err != nil { fmt.Println("binary.Write failed:", err)}
//
//
//fmt.Println(buf.Len())
//
//fmt.Printf("%x", buf.Bytes())
//
//fmt.Println(buf.Bytes())

