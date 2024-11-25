package model

type Customer struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size:170;unique"`
	Phone     string    `json:"phone" gorm:"size:30"`
	Email     string    `json:"email" gorm:"size:170"`
	Document  string    `json:"document" gorm:"size:35"`
	Address   string    `json:"address" gorm:"size:255"`
	Contacts  []Contact `json:"contacts" gorm:"serializer:json"`
	Active    bool      `json:"active"`
	CreatedAt int64     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64     `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt int64     `json:"deleted_at" gorm:"default:null"`
}

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
