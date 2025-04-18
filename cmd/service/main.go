package main

import (
	"fmt"

	"github.com/pondparinya/go-structure/internal/config"
)

func main() {
	var cfg config.Config

	err := config.Load("configs", "config", "yaml", &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded config: %+v\n", cfg)
}
