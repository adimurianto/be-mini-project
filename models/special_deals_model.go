package models

type SpecialDeals struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	SpecialDealsBase
}

type SpecialDealsBase struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Discount int    `json:"discount"`
	Image    string `json:"image"`
	Status   bool   `json:"status"`
}

// TableName is Database TableName of this model
func (e *SpecialDeals) TableName() string {
	return "special_deals"
}
