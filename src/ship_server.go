package main

import (
	"chaos"
)

func main() {
	tartar := chaos.Tartar{Port:":9009"}
	tartar.Run()
}
