package main

import (
	"bloger_server/core"
	"bloger_server/flags"
	"bloger_server/global"
	"fmt"

	_ "github.com/sirupsen/logrus"
)

func main() {

	flags.Parse()

	fmt.Println(flags.FlagOptions)

	global.Config = core.ReadConf()
	global.DB = core.InitDB()
	core.InitLogrus()

	core.InitIPDB()
}
