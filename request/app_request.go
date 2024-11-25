package request

import "optimus.dev.br/gaia/model"

type AppCreateRequest struct {
	Name       string          `json:"name"`
	Image      string          `json:"image"`
	Env        []string        `json:"env"`
	Cmd        []string        `json:"cmd"`
	Replicas   int             `json:"replicas"`
	Listening  bool            `json:"listening"`
	Ports      []model.AppPort `json:"ports"`
	HostPort   int             `json:"host_port"`
	Protocol   string          `json:"protocol"`
	Active     bool            `json:"active"`
	CustomerID int64           `json:"customer_id"`
}
