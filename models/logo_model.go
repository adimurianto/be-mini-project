package models

type Logo struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	LogoBase
}

type LogoBase struct {
	Title  string `json:"title"`
	Logo   string `json:"logo"`
	Status bool   `json:"status"`
}

// TableName is Database TableName of this model
func (e *Logo) TableName() string {
	return "logo"
}
