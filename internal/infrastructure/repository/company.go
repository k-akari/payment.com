package repository

import (
	"context"
	"fmt"

	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
)

type companyRepository struct {
	dba database.Client
}

func NewCompanyRepository(
	dba database.Client,
) *companyRepository {
	return &companyRepository{
		dba: dba,
	}
}

func (r *companyRepository) Create(
	ctx context.Context,
	company *domain.Company,
) error {
	const query = `INSERT INTO companies (name, representative, telephone_number, postal_code, address) VALUES (:name, :representative, :telephone_number, :postal_code, :address)`
	result, err := r.dba.Exec(ctx, query, *company)
	if err != nil {
		return fmt.Errorf("failed to run r.dba.Exec: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to run result.RowsAffected: %w", err)
	}

	return nil
}

func (r *companyRepository) GetByID(
	ctx context.Context,
	companyID domain.CompanyID,
) (*domain.Company, error) {
	const q = `SELECT * FROM companies where id=:company_id;`
	a := struct {
		CompanyID uint64 `db:"company_id"`
	}{
		CompanyID: uint64(companyID),
	}
	rows, err := r.dba.Query(ctx, q, a)
	if err != nil {
		return nil, fmt.Errorf("failed to run r.dba.Query: %w", err)
	}
	defer rows.Close()

	if ok := rows.Next(); !ok {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error during rows iteration: %w", err)
		}
		return nil, fmt.Errorf("not found company: %q", companyID)
	}

	var company domain.Company
	if err := rows.StructScan(&company); err != nil {
		return nil, fmt.Errorf("failed to run rows.StructScan: %w", err)
	}

	return &company, nil
}
