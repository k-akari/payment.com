package domain

import (
	"time"

	"go.mercari.io/go-bps/bps"
)

type (
	InvoiceID     uint64
	InvoiceStatus uint8

	Invoice struct {
		ID             InvoiceID     `db:"id"`
		CompanyID      CompanyID     `db:"company_id"`
		ClientID       ClientID      `db:"client_id"`
		IssuedDate     *time.Time    `db:"issued_date"`
		PaidAmount     int32         `db:"paid_amount"`
		Fee            int32         `db:"fee"`
		FeeRate        *bps.BPS      `db:"fee_rate"`
		SalesTax       int32         `db:"sales_tax"`
		SalesTaxRate   *bps.BPS      `db:"sales_tax_rate"`
		BilledAmount   int32         `db:"billed_amount"`
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
