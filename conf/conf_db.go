package conf

import "fmt"

type DB struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	DBName       string `yaml:"dbname"`
	Timeout      string `yaml:"timeout"`
	Debug        bool   `yaml:"debug"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func (d DB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", d.Username, d.Password, d.Host, d.Port, d.DBName, d.Timeout)
}
