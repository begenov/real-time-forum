package main

import (
	"log"

	"github.com/begenov/real-time-forum/internal/app"
	"github.com/begenov/real-time-forum/internal/config"
)

const path = "./internal/config/config.json"

func main() {
	cfg, err := config.Init(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}

}
