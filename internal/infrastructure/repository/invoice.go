package repository

import (
	"context"
	"fmt"

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
