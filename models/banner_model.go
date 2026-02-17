package models

type Banner struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	BannerBase
}

type BannerBase struct {
	Title          string `json:"title"`
	Link           string `json:"link"`
	PrimaryImage   string `json:"primary_image"`
	SecondaryImage string `json:"secondary_image"`
	Status         bool   `json:"status"`
}

// TableName is Database TableName of this model
func (e *Banner) TableName() string {
	return "banner"
}
