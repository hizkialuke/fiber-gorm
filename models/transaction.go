package models

import "time"

type Transaction struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	MerchantId int64     `json:"merchant_id"`
	OutletId   int64     `json:"outlet_id"`
	BillTotal  float64   `json:"bill_total"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int64     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int64     `json:"updated_by"`
}
