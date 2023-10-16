package domain

import "time"

type (
	ClientID uint64

	Client struct {
		ID              ClientID   `db:"id" json:"id"`
		CompanyID       CompanyID  `db:"company_id" json:"company_id"`
		Name            string     `db:"name" json:"name"`
		Representative  string     `db:"representative" json:"representative"`
		TelephoneNumber string     `db:"telephone_number" json:"telephone_number"`
		PostalCode      string     `db:"postal_code" json:"postal_code"`
		Address         string     `db:"address" json:"address"`
		CreatedAt       *time.Time `db:"created_at" json:"created_at"`
		UpdatedAt       *time.Time `db:"updated_at" json:"updated_at"`
	}
)
