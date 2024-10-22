package main

import (
	"web3/config"
	"web3/model"
	"web3/route"
	"web3/utils"
)

func main() {
	config.ViperInit()
	model.ConnectMysql()
	model.InitRedis()
	route.Init()
	utils.InitListen()
}
