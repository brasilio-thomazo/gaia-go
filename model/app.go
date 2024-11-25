package model

import "github.com/google/uuid"

type App struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;"`
	CustomerID  int64     `json:"customer_id"`
	Name        string    `json:"name" gorm:"unique;size:50"`
	ContainerID string    `json:"container_id" gorm:"unique;size:255"`
	Image       string    `json:"image" gorm:"size:255"`
	Env         []string  `json:"env" gorm:"serializer:json"`
	Cmd         []string  `json:"cmd" gorm:"serializer:json"`
	Ports       []AppPort `json:"ports" gorm:"type:json;serializer:json"`
	Replicas    int       `json:"replicas" gorm:"default:1"`
	Listening   bool      `json:"listening"`
	Active      bool      `json:"active"`
	Customer    Customer  `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   int64     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   int64     `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt   int64     `json:"deleted_at" gorm:"default:null"`
}

type AppPort struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	HostPort string `json:"host_port"`
}
