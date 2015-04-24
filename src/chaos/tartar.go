package chaos

import (
	"fmt"
	"net"
	"os"
	//"time"
	"binary"
	m "messages"
	"messages"
)

type Tartar struct {
	Port string
	conn *net.UDPConn
	gaia *Gaia
}

func (t *Tartar) Run() {
	if t.initiatedUDP() {
		fmt.Println("Tartar is run")

		t.gaia = &Gaia{tickDuration: 100}
		t.gaia.Run()

		for {
			t.handleClient()
		}
	}
}

func (t *Tartar) handleClient() {
	var buf [512]byte
	readCount, addr, err := t.conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error [handleClient]:", err.Error())
		return
	} else {
		fmt.Printf("%v:%v (%v)\n", addr.IP, addr.Port, addr.Zone)
		fmt.Printf("[%v] %X\n", readCount, buf[0:readCount])
		t.processBuffer(buf[0:], addr)
	}
}

func (c *Tartar) initiatedUDP() bool {
	var udpAddr *net.UDPAddr
	var err error

	udpAddr, err = net.ResolveUDPAddr("udp", c.Port)
	if err != nil {
		fmt.Printf("ResolveUDPAddr: %v\n", err.Error())
		return false
	}

	c.conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("ListenUDP: %v\n", err.Error())
		return false
	}
	return true
}

func (t *Tartar) processBuffer(bytes []byte, addr *net.UDPAddr) {
	messageCode := binary.ReadUInt16(bytes)

	switch messageCode {
	case m.LOGIN_REQUEST_CODE:
		msg := messages.GetLoginRequestFromBytes(bytes[2:])
		pilot, ok := pilots[msg.GetLogin()]
		if !ok {
			AddPilot(msg.GetLogin(), addr)
			fmt.Printf("Added pilot: %v\n", msg.GetLogin())
		}else{
			fmt.Printf("!!! Second login from: %v %v %v\n", msg.GetLogin(),pilot.addr,addr)
			pilot.addr = addr
		}
	case m.SHIP_DRIVING_CODE:
		msg := messages.GetShipDrivingFromBytes(bytes[2:])
		pilot, ok := pilots[msg.GetLogin()]
		if ok {
			pilot.touch()
			t.gaia.queue <- msg
		}else{
			fmt.Printf("!!! Ship driving whithout login: %v\n", msg.GetLogin())
		}
	case m.ADD_SHIP_CODE:
		msg := messages.GetAddShipFromBytes(bytes[2:])
		pilot, ok := pilots[msg.GetLogin()]
		if ok {
			pilot.touch()
			t.gaia.queue <- msg
		}else{
			fmt.Printf("!!! Ship adding whithout login: %v\n", msg.GetLogin())
		}
	default:
		fmt.Printf("!!! Unknown message code: %v\n", messageCode)
	}
}
