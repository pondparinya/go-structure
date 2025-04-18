package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Load(path, name, ext string, cfg any) error {
	v := viper.New()

	v.SetConfigName(name)
	v.SetConfigType(ext)
	v.AddConfigPath(path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err := v.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode config into struct: %w", err)
	}

	// Watch for changes and hot-reload
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)

		if err := v.Unmarshal(cfg); err != nil {
			log.Printf("Error unmarshaling config: %v", err)
		} else {
			log.Println("Config reloaded successfully")
		}
	})
	return nil
}
