package domain

import "time"

type (
	ClientBankAccountID uint64

	ClientBankAccount struct {
		ID            ClientBankAccountID `db:"id" json:"id"`
		ClientID      ClientID            `db:"client_id" json:"client_id"`
		BankName      string              `db:"bank_name" json:"bank_name"`
		BranchName    string              `db:"branch_name" json:"branch_name"`
		AccountNumber string              `db:"account_number" json:"account_number"`
		AccountName   string              `db:"account_name" json:"account_name"`
		CreatedAt     time.Time           `db:"created_at" json:"created_at"`
		UpdatedAt     time.Time           `db:"updated_at" json:"updated_at"`
	}
)
