package model

type User struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID   int    `json:"group_id"`
	Name      string `json:"name" gorm:"size:100"`
	Phone     string `json:"phone" gorm:"size:30"`
	JobTitle  string `json:"job_title" gorm:"size:50"`
	Email     string `json:"email" gorm:"unique;size:170"`
	Username  string `json:"username" gorm:"unique;size:50"`
	Password  string `json:"-" gorm:"size:255"`
	Visible   bool   `json:"visible"`
	Editable  bool   `json:"editable"`
	Locked    bool   `json:"locked"`
	Group     Group  `json:"group" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt int64  `json:"deleted_at" gorm:"default:null"`
}
