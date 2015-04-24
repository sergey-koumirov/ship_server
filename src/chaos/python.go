package chaos

import(
    "fmt"
	"net"
)

type Python struct{
	message chan []byte
	conn *net.UDPConn
}


func (p *Python) Run(){
	p.message = make(chan []byte, 1000)

	var err error
	p.conn, err = net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9008})
    chk(err)

    go p.tick()
	fmt.Println("Python is run")
}

func (p *Python) tick() {

	for {
		bytes := <-p.message
		//fmt.Println("Bytes to send")
		//fmt.Println(bytes)

		SendBroadcast(bytes, p.conn)
	}

}


