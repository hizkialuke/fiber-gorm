package models

import "time"

type Merchant struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	UserId       int64     `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int64     `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    int64     `json:"updated_by"`
}
