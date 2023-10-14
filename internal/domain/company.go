package domain

import "time"

type (
	CompanyID uint64

	Company struct {
		ID              CompanyID  `db:"id"`
		Name            string     `db:"name"`
		Representative  string     `db:"representative"`
		TelephoneNumber string     `db:"telephone_number"`
		PostalCode      string     `db:"postal_code"`
		Address         string     `db:"address"`
		CreatedAt       *time.Time `db:"created_at"`
		UpdatedAt       *time.Time `db:"updated_at"`
	}
)
