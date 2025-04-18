package config

import "time"

type (
	Config struct {
		MongoDB struct {
			URI       string
			Addresses []string
			Username  string
			Password  string
			Timeout   time.Duration
		}
	}
)
