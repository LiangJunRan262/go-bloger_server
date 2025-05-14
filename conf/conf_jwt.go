package conf

type JWT struct {
	Secret string `yaml:"secret"`
	Expire int64  `yaml:"expire"`
	Issuer string `yaml:"issuer"`
}
