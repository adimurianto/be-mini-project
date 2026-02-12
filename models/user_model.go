package models

import "time"

type User struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	UserBase
	CreatedAt time.Time `json:"created_at"`
}

type UserBase struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
}

// TableName is Database TableName of this model
func (e *User) TableName() string {
	return "user"
}
