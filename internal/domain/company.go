package domain

import "time"

type (
	CompanyID uint64

	Company struct {
		ID              CompanyID `db:"id" json:"id"`
		Name            string    `db:"name" json:"name"`
		Representative  string    `db:"representative" json:"representative"`
		TelephoneNumber string    `db:"telephone_number" json:"telephone_number"`
		PostalCode      string    `db:"postal_code" json:"postal_code"`
		Address         string    `db:"address" json:"address"`
		CreatedAt       time.Time `db:"created_at" json:"created_at"`
		UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
	}
)
