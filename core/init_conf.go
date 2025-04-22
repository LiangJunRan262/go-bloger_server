package core

import (
	"bloger_server/conf"
	"bloger_server/flags"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// var confPath string = "settings.yaml" // confPath - путь к файлу конфигурации

func ReadConf() (c *conf.Config) {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err) // TODO: обработать ошибку чтения конфигурации
	}

	c = new(conf.Config)

	// var conf Config

	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误 %s", err)) // TODO: обработать ошибку парсинга конфигурации
	}

	fmt.Println(c)
	fmt.Println(c.System.Ip)
	fmt.Println(c.System.Port)

	return
}
