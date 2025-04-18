package config

import "time"

type (
	Config struct {
		Stage   string  `env:"STAGE"`
		Port    string  `mapstructure:"port"`
		MongoDB MongoDB `mapstructure:"mongodb"`
	}
	MongoDB struct {
		URI       string        `mapstructure:"uri"`
		Addresses []string      `mapstructure:"addresses"`
		Username  string        `mapstructure:"username"`
		Password  string        `mapstructure:"password"`
		Timeout   time.Duration `mapstructure:"timeout"`
	}
)
