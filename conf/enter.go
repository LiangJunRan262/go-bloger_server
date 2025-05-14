package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`
	JWT    JWT    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
}
