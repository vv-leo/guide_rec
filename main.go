package main

import (
	"web3/config"
	"web3/model"
	"web3/route"
)

func main() {

	config.ViperInit()

	model.InitMysql()

	route.Init()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
