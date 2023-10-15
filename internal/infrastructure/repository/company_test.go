package repository

import (
	"context"
	"testing"

	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
	"github.com/k-akari/payment.com/internal/testutil"
)

type companyRepositoryTestHelper struct {
	sub *companyRepository
}

func createCompanyRepositoryTestHelper(t *testing.T) *companyRepositoryTestHelper {
	sqlxDB, dbCloseFunc := testutil.InitDB()
	dba := database.NewClient(sqlxDB)

	t.Cleanup(dbCloseFunc)

	return &companyRepositoryTestHelper{
		sub: NewCompanyRepository(dba),
	}
}

func TestCompanyRepository_Create(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		h := createCompanyRepositoryTestHelper(t)

		company := domain.Company{
			Name:            "name",
			Representative:  "representative",
			TelephoneNumber: "telephone_number",
			PostalCode:      "postal_code",
			Address:         "address",
		}
		err := h.sub.Create(context.Background(), &company)
		if err != nil {
			t.Errorf("err should be nil, but got: %v", err)
		}
	})
}
