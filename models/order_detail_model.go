package models

type OrderDetail struct {
	ID string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	OrderDetailBase
}

type OrderDetailBase struct {
	OrderID   string `gorm:"type:uuid;not null;index" json:"order_id"`
	ProductID string `gorm:"type:text;not null" json:"product_id"`
	Quantity  int64  `gorm:"type:bigint;not null" json:"quantity"`
	Price     int64  `gorm:"type:bigint;not null" json:"price"`
	Status    bool   `gorm:"default:true" json:"status"`
}

type OrderDetailResponse struct {
	OrderDetail
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

type OrderDetailRequest struct {
	ProductID string `gorm:"type:text;not null" json:"product_id"`
	Quantity  int64  `gorm:"type:bigint;not null" json:"quantity"`
	Price     int64  `gorm:"type:bigint;not null" json:"price"`
}

// TableName is Database TableName of this model
func (e *OrderDetail) TableName() string {
	return "order_detail"
}

func (e *OrderDetailResponse) TableName() string {
	return "order_detail"
}
