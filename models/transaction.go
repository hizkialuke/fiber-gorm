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

	Merchant *Merchant
	Outlet   *Outlet
}

type ReportRequest struct {
	StartDate *string `query:"start_date"`
	EndDate   *string `query:"end_date"`
	*Pagination
}

type ReportATransformer struct {
	Date         string  `json:"date"`
	MerchantName string  `json:"merchant_name"`
	Omzet        float64 `json:"omzet"`
}

func (t *Transaction) ToReportA() *ReportATransformer {
	return &ReportATransformer{
		Date:         t.CreatedAt.Format("2006-01-02"),
		MerchantName: t.Merchant.MerchantName,
		Omzet:        t.BillTotal,
	}
}

type ReportBTransformer struct {
	Date         string  `json:"date"`
	MerchantName string  `json:"merchant_name"`
	OutletName   string  `json:"outlet_name"`
	Omzet        float64 `json:"omzet"`
}

func (t *Transaction) ToReportB() *ReportBTransformer {
	return &ReportBTransformer{
		Date:         t.CreatedAt.Format("2006-01-02"),
		MerchantName: t.Outlet.Merchant.MerchantName,
		OutletName:   t.Outlet.OutletName,
		Omzet:        t.BillTotal,
	}
}
