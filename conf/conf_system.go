package conf

type System struct {
	Name string `yaml:"name"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
