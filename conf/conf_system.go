package conf

import "fmt"

type System struct {
	Name    string `yaml:"name"`
	Ip      string `yaml:"ip"`
	Port    int    `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
}

func (s System) GetAddr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
