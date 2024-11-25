package model

type Group struct {
	ID          int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string   `json:"name" gorm:"unique;size:50"`
	Permissions []string `json:"permissions" gorm:"serializer:json"`
	Visible     bool     `json:"visible"`
	Editable    bool     `json:"editable"`
	Locked      bool     `json:"locked"`
	CreatedAt   int64    `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   int64    `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt   int64    `json:"deleted_at" gorm:"default:null"`
}
