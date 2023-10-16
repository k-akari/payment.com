package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/k-akari/payment.com/internal/domain"
)

type invoiceRepository struct {
	dba accessor
}

func NewInvoiceRepository(
	dba accessor,
) *invoiceRepository {
	return &invoiceRepository{
		dba: dba,
	}
}

func (r *invoiceRepository) Create(
	ctx context.Context,
	invoice *domain.Invoice,
) (domain.InvoiceID, error) {
	const query = `INSERT INTO invoices (company_id, client_id, issued_date, paid_amount, fee, fee_rate, sales_tax, sales_tax_rate, billed_amount, payment_due_date, status) VALUES (:company_id, :client_id, :issued_date, :paid_amount, :fee, :fee_rate, :sales_tax, :sales_tax_rate, :billed_amount, :payment_due_date, :status)`
	result, err := r.dba.Exec(ctx, query, *invoice)
	if err != nil {
		return domain.InvoiceID(0), fmt.Errorf("failed to run r.dba.Exec: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.InvoiceID(0), fmt.Errorf("failed to run result.RowsAffected: %w", err)
	}

	iid, err := result.LastInsertId()
	if err != nil {
		return domain.InvoiceID(0), fmt.Errorf("failed to run result.LastInsertId: %w", err)
	}

	return domain.InvoiceID(iid), nil
}

func (r *invoiceRepository) ListByPaymentDueDateBetween(
	ctx context.Context,
	coid domain.CompanyID,
	from *time.Time,
	to *time.Time,
) ([]*domain.Invoice, error) {
	const q = `SELECT * FROM invoices WHERE company_id=:company_id AND payment_due_date BETWEEN :from AND :to;`
	a := struct {
		CompanyID domain.CompanyID `db:"company_id"`
		From      *time.Time       `db:"from"`
		To        *time.Time       `db:"to"`
	}{
		CompanyID: coid,
		From:      from,
		To:        to,
	}
	rows, err := r.dba.Query(ctx, q, a)
	if err != nil {
		return nil, fmt.Errorf("failed to run r.dba.Query: %w", err)
	}
	defer rows.Close()

	var invoices []*domain.Invoice
	for rows.Next() {
		var ir domain.InvoiceRecord
		err := rows.StructScan(&ir)
		if err != nil {
			return nil, fmt.Errorf("failed to run rows.StructScan: %w", err)
		}
		invoice, err := ir.ConvertIntoInvoice()
		if err != nil {
			return nil, fmt.Errorf("failed to run ir.ConvertInvoice: %w", err)
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}
