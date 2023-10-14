package domain

import "time"

type (
	UserID uint64

	User struct {
		ID        UserID     `db:"id"`
		CompanyID CompanyID  `db:"company_id"`
		Name      string     `db:"name"`
		Email     string     `db:"email"`
		Password  string     `db:"password"`
		CreatedAt *time.Time `db:"created_at"`
		UpdatedAt *time.Time `db:"updated_at"`
	}
)
