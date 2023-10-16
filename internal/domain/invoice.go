package domain

import (
	"time"

	"go.mercari.io/go-bps/bps"
)

const (
	defaultFeeRate      = 4
	defaultSalesTaxRate = 10
)

type (
	InvoiceID     uint64
	InvoiceStatus uint8

	Invoice struct {
		ID             InvoiceID     `db:"id"`
		CompanyID      CompanyID     `db:"company_id"`
		ClientID       ClientID      `db:"client_id"`
		IssuedDate     *time.Time    `db:"issued_date"`
		PaidAmount     int64         `db:"paid_amount"`
		Fee            int64         `db:"fee"`
		FeeRate        *bps.BPS      `db:"fee_rate"`
		SalesTax       int64         `db:"sales_tax"`
		SalesTaxRate   *bps.BPS      `db:"sales_tax_rate"`
		BilledAmount   int64         `db:"billed_amount"`
		PaymentDueDate *time.Time    `db:"payment_due_date"`
		Status         InvoiceStatus `db:"status"`
		CreatedAt      *time.Time    `db:"created_at"`
		UpdatedAt      *time.Time    `db:"updated_at"`
	}
)

const (
	InvoiceStatusUnpaid InvoiceStatus = iota
	InvoiceStatusProcessing
	InvoiceStatusPaid
	InvoiceStatusError
)

func (i *Invoice) SetDefaultRate() {
	i.FeeRate = bps.NewFromPercentage(defaultFeeRate)
	i.SalesTaxRate = bps.NewFromPercentage(defaultSalesTaxRate)
}

func (i *Invoice) CalcBilledAmount() {
	fee := i.FeeRate.Mul(i.PaidAmount)
	salesTax := i.SalesTaxRate.Mul(i.PaidAmount)

	i.Fee = fee.Amounts()
	i.SalesTax = salesTax.Amounts()
	i.BilledAmount = bps.Sum(bps.NewFromAmount(i.PaidAmount), fee, salesTax).Amounts()
}
