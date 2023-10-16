package domain

import "time"

type (
	UserID uint64

	User struct {
		ID        UserID     `db:"id" json:"id"`
		CompanyID CompanyID  `db:"company_id" json:"company_id"`
		Name      string     `db:"name" json:"name"`
		Email     string     `db:"email" json:"email"`
		Password  string     `db:"password" json:"password"`
		CreatedAt *time.Time `db:"created_at" json:"created_at"`
		UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	}
)
