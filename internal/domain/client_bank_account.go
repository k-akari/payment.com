package domain

import "time"

type (
	ClientBankAccountID uint64

	ClientBankAccount struct {
		ID            ClientBankAccountID `db:"id"`
		ClientID      ClientID            `db:"client_id"`
		BankName      string              `db:"bank_name"`
		BranchName    string              `db:"branch_name"`
		AccountNumber string              `db:"account_number"`
		AccountName   string              `db:"account_name"`
		CreatedAt     *time.Time          `db:"created_at"`
		UpdatedAt     *time.Time          `db:"updated_at"`
	}
)
