package conf

import "github.com/jasontconnell/conf"

type Config struct {
	ConnectionString string `json:"connectionString"`
}

func LoadConfig(configFilename string) (Config, error) {
	var cfg Config
	err := conf.LoadConfig(configFilename, &cfg)
	return cfg, err
}
