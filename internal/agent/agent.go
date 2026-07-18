package app

import (
	"github.com/Alexunder2003/alex-metrics-service/internal/config"
	service "github.com/Alexunder2003/alex-metrics-service/internal/service/agent"
)

type Agent struct {
	agentService *service.AgentService
}

func NewAgent(cfg *config.Config) *Agent {
	return &Agent{agentService: service.NewAgentService(cfg)}
}

func (a *Agent) Run() {
	a.agentService.Run()
}