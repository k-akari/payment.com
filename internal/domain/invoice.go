package domain

import (
	"fmt"
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
		ID             InvoiceID     `db:"id" json:"id"`
		CompanyID      CompanyID     `db:"company_id" json:"company_id"`
		ClientID       ClientID      `db:"client_id" json:"client_id"`
		IssuedDate     time.Time     `db:"issued_date" json:"issued_date"`
		PaidAmount     int64         `db:"paid_amount" json:"paid_amount"`
		Fee            int64         `db:"fee" json:"fee"`
		FeeRate        *bps.BPS      `db:"fee_rate" json:"fee_rate"`
		SalesTax       int64         `db:"sales_tax" json:"sales_tax"`
		SalesTaxRate   *bps.BPS      `db:"sales_tax_rate" json:"sales_tax_rate"`
		BilledAmount   int64         `db:"billed_amount" json:"billed_amount"`
		PaymentDueDate time.Time     `db:"payment_due_date" json:"payment_due_date"`
		Status         InvoiceStatus `db:"status" json:"status"`
		CreatedAt      time.Time     `db:"created_at" json:"created_at"`
		UpdatedAt      time.Time     `db:"updated_at" json:"updated_at"`
	}

	InvoiceRecord struct {
		ID             InvoiceID     `db:"id"`
		CompanyID      CompanyID     `db:"company_id"`
		ClientID       ClientID      `db:"client_id"`
		IssuedDate     time.Time     `db:"issued_date"`
		PaidAmount     int64         `db:"paid_amount"`
		Fee            int64         `db:"fee"`
		FeeRate        string        `db:"fee_rate"`
		SalesTax       int64         `db:"sales_tax"`
		SalesTaxRate   string        `db:"sales_tax_rate"`
		BilledAmount   int64         `db:"billed_amount"`
		PaymentDueDate time.Time     `db:"payment_due_date"`
		Status         InvoiceStatus `db:"status"`
		CreatedAt      time.Time     `db:"created_at"`
		UpdatedAt      time.Time     `db:"updated_at"`
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

func (ir *InvoiceRecord) ConvertIntoInvoice() (*Invoice, error) {
	feeRate, err := bps.NewFromString(ir.FeeRate)
	if err != nil {
		return nil, fmt.Errorf("failed to run bps.NewFromString: %w", err)
	}
	salesTaxRate, err := bps.NewFromString(ir.SalesTaxRate)
	if err != nil {
		return nil, fmt.Errorf("failed to run bps.NewFromString: %w", err)
	}

	return &Invoice{
		ID:             ir.ID,
		CompanyID:      ir.CompanyID,
		ClientID:       ir.ClientID,
		IssuedDate:     ir.IssuedDate,
		PaidAmount:     ir.PaidAmount,
		Fee:            ir.Fee,
		FeeRate:        feeRate,
		SalesTax:       ir.SalesTax,
		SalesTaxRate:   salesTaxRate,
		BilledAmount:   ir.BilledAmount,
		PaymentDueDate: ir.PaymentDueDate,
		Status:         ir.Status,
		CreatedAt:      ir.CreatedAt,
		UpdatedAt:      ir.UpdatedAt,
	}, nil
}
