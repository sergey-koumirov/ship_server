package chaos

import(
	"net"
	"time"
	"sync"
	//"fmt"
)

type Locker struct{
	sync.RWMutex
}

var (
	pilots = map[string]*Pilot{}
	locker = new(Locker)
)


type Pilot struct {
	login string
	addr *net.UDPAddr
	loginTime time.Time
	lastMessageTime time.Time
}

func AddPilot(login string, addr *net.UDPAddr){
	locker.RLock()
	t := time.Now()
	pilots[login] = &Pilot{login: login, addr: addr, loginTime: t, lastMessageTime: t}
	locker.RUnlock()
}

func (p *Pilot) touch(){
	p.lastMessageTime = time.Now()
}

func SendBroadcast(bytes []byte, conn *net.UDPConn){
	locker.RLock()

	for _, p := range pilots {
		//fmt.Println(p.addr)
		//fmt.Println(bytes)

		_, err := conn.WriteToUDP(bytes, p.addr)
		chk(err)
	}

	locker.RUnlock()
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

