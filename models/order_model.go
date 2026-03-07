package models

import "time"

type Order struct {
	ID string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	OrderBase
}

type OrderBase struct {
	Table         string    `gorm:"type:text;not null" json:"table"`
	Fullname      string    `gorm:"type:text;not null" json:"fullname"`
	Phone         string    `gorm:"type:text" json:"phone"`
	IpAddress     string    `gorm:"type:text;column:ip_address" json:"ip_address"`
	TotalPrice    int64     `gorm:"type:bigint;not null" json:"total_price"`
	StatusPayment bool      `gorm:"column:status_payment;default:false" json:"status_payment"`
	Status        bool      `gorm:"default:true" json:"status"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
}

type OrderResponse struct {
	Order
	OrderDetails []OrderDetailResponse `gorm:"foreignKey:OrderID" json:"order_details,omitempty"`
}

type OrderRequest struct {
	Table      string `gorm:"type:text;not null" json:"table"`
	Fullname   string `gorm:"type:text;not null" json:"fullname"`
	Phone      string `gorm:"type:text" json:"phone"`
	IpAddress  string `gorm:"type:text;column:ip_address" json:"ip_address"`
	TotalPrice int64  `gorm:"type:bigint;not null" json:"total_price"`

	OrderDetails []OrderDetailRequest `gorm:"foreignKey:OrderID" json:"order_details,omitempty"`
}

// TableName is Database TableName of this model
func (e *Order) TableName() string {
	return "order"
}

func (e *OrderResponse) TableName() string {
	return "order"
}
