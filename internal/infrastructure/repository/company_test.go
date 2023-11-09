package repository

import (
	"context"
	"testing"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/infrastructure/database"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil"
)

type companyRepositoryTestHelper struct {
	sub *CompanyRepository
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
		cid, err := h.sub.Create(context.Background(), &company)
		if err != nil {
			t.Errorf("err should be nil, but got: %v", err)
		}
		if cid == 0 {
			t.Errorf("domain.CompanyID should not be 0")
		}
	})
}
