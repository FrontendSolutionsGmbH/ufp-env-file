package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	Prefix     string        `env:"PREFIX"      envDefault:"UFP"`
	Separator  string        `env:"SEPARATOR"   envDefault:"-"`
	TargetFile string        `env:"TARGET_FILE" envDefault:"/output.json"`
	FileType   string        `env:"FILE_TYPE"   envDefault:"JSON"`
}

func GetConfig() Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", cfg)
	return cfg
}
