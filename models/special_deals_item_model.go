package models

type SpecialDealsItem struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	SpecialDealsItemBase
}

type SpecialDealsItemBase struct {
	ProductID      string `json:"product_id"`
	SpecialDealsID string `json:"special_deals_id"`
	Quantity       int    `json:"quantity"`
	Status         bool   `json:"status"`
}

// TableName is Database TableName of this model
func (e *SpecialDealsItem) TableName() string {
	return "special_deals_item"
}
