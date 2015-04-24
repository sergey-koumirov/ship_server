package chaos

import (
	"abyss"
	"fmt"
	"messages"
	"time"
	//"reflect"
)

type Gaia struct {
	tickDuration int32
	ticker       *time.Ticker
	world        *abyss.WorldImage
	python       *Python
	queue        chan messages.Message
}

func (g *Gaia) Run() {
	g.ticker = time.NewTicker(time.Millisecond * time.Duration(g.tickDuration))
	g.world = new(abyss.WorldImage)
	g.world.LoadShipTemplates()
	g.queue = make(chan messages.Message, 1000)
	g.python = &Python{}
	g.python.Run()

	fmt.Println("Gaia is run")
	go g.tick()
}

func (g *Gaia) tick() {
	lastT := time.Now()
	for t := range g.ticker.C {
		g.world.TickTime = t

		select {
		case m := <-g.queue:
			g.applyMessage(m)
		default:
			//fmt.Println("Gaia queue is emptyed")
		}

		g.world.MoveObjects( t.Sub(lastT).Seconds() )

		bytes := g.world.ToBytes()
		g.python.message <- bytes

		//fmt.Println("\n---------------------------------")
		//fmt.Println( t.Sub(lastT).Seconds() )

		lastT = t
	}

}

func (g *Gaia) applyMessage(m messages.Message) {
	switch interface{}(m).(type) {
	case *messages.AddShip:
		g.world.AddShip( m.GetLogin() )
	case *messages.ShipDriving:
		g.world.UpdateShipControl( m.(*messages.ShipDriving) )
	default:
		fmt.Println("Unknown message type.")
	}
}
