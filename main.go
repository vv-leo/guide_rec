package main

import (
	"web3/config"
	"web3/ether/listen"
	"web3/model"
	"web3/route"
)

func main() {
	config.ViperInit()
	model.ConnectMysql()
	model.InitRedis()
	listen.EventListen()
	route.Init()
}
