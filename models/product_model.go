package models

type Product struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	ProductBase
}

type ProductBase struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	CategoryID  string `json:"category_id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

// TableName is Database TableName of this model
func (e *Product) TableName() string {
	return "product"
}
