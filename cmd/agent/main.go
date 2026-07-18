package main

import (
	"github.com/Alexunder2003/alex-metrics-service/internal/agent"
	"github.com/Alexunder2003/alex-metrics-service/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	agent := agent.NewAgent(&cfg)
	agent.Run()
}