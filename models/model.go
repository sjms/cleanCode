package models

import (
	"github.com/google/uuid"
)

type Customers struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id,omitempty"`
	Name string    `uri:"name" binding:"required"`
}

type Products struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id,omitempty"`
	Description string    `gorm:"not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Quantity    int       `gorm:"not null" json:"quantity"`
}

type Orders struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id,omitempty"`
	ProductId  uuid.UUID `gorm:"type:uuid;references:Id" json:"product_id"`
	CouponCode string    `gorm:"not null" json:"coupon_code,omitempty"`
}
