package main

import (
	"gin_consult/config"
	"gin_consult/routers"
)

func main() {
	// Ek1+Ep1==Ek2+Ep2
	config.Init()
	r := routers.NewRouter()
	_=r.Run(config.HttpPort)
}