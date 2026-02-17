package models

type Category struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	CategoryBase
}

type CategoryBase struct {
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Status bool   `json:"status"`
}

// TableName is Database TableName of this model
func (e *Category) TableName() string {
	return "category"
}
