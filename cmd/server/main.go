package main

import (
	"log"

	"github.com/Alexunder2003/alex-metrics-service/internal/app"
	"github.com/Alexunder2003/alex-metrics-service/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.NewApp(cfg).Run(); err != nil {
		log.Fatal(err)
	}
}
