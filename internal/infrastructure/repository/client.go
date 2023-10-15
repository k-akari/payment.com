package repository

import (
	"context"
	"fmt"

	"github.com/k-akari/payment.com/internal/domain"
)

type clientRepository struct {
	dba accessor
}

func NewClientRepository(
	dba accessor,
) *clientRepository {
	return &clientRepository{
		dba: dba,
	}
}

func (r *clientRepository) Create(
	ctx context.Context,
	client *domain.Client,
) (domain.ClientID, error) {
	const query = `INSERT INTO clients (company_id, name, representative, telephone_number, postal_code, address) VALUES (:company_id, :name, :representative, :telephone_number, :postal_code, :address)`
	result, err := r.dba.Exec(ctx, query, *client)
	if err != nil {
		return domain.ClientID(0), fmt.Errorf("failed to run r.dba.Exec: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.ClientID(0), fmt.Errorf("failed to run result.RowsAffected: %w", err)
	}

	cid, err := result.LastInsertId()
	if err != nil {
		return domain.ClientID(0), fmt.Errorf("failed to run result.LastInsertId: %w", err)
	}

	return domain.ClientID(cid), nil
}
