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

func (r *clientRepository) GetByID(
	ctx context.Context,
	clientID domain.ClientID,
) (*domain.Client, error) {
	const q = `SELECT * FROM clients where id=:client_id;`
	a := struct {
		ClientID uint64 `db:"client_id"`
	}{
		ClientID: uint64(clientID),
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
		return nil, fmt.Errorf("not found company: %q", clientID)
	}

	var client domain.Client
	if err := rows.StructScan(&client); err != nil {
		return nil, fmt.Errorf("failed to run rows.StructScan: %w", err)
	}

	return &client, nil
}
