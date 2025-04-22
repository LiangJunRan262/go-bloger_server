package flags

import "flag"

type Options struct {
	File    string
	DB      bool
	Version string
}

var FlagOptions = new(Options) // FlagOptions - глобальная переменная для хранения флагов командной строки

func Parse() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.StringVar(&FlagOptions.Version, "v", "0.0.1", "版本号")
	flag.Parse()
}
