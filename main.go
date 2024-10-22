package main

import (
	"web3/config"
	"web3/model"
	"web3/route"
)

func main() {
	config.ViperInit()
	model.ConnectMysql()
	model.InitRedis()
	//utils.InitListen()
	route.Init()
}
