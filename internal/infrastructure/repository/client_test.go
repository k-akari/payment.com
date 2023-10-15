package repository

import (
	"context"
	"testing"

	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
	"github.com/k-akari/payment.com/internal/testutil"
)

type clientRepositoryTestHelper struct {
	sub               *clientRepository
	companyRepository *companyRepository
}

func createClientRepositoryTestHelper(t *testing.T) *clientRepositoryTestHelper {
	sqlxDB, dbCloseFunc := testutil.InitDB()
	dba := database.NewClient(sqlxDB)

	t.Cleanup(dbCloseFunc)

	return &clientRepositoryTestHelper{
		sub:               NewClientRepository(dba),
		companyRepository: NewCompanyRepository(dba),
	}
}

func TestClientRepository_Create(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		h := createClientRepositoryTestHelper(t)

		company := domain.Company{
			Name:            "company_name",
			Representative:  "company_representative",
			TelephoneNumber: "company_telephone_number",
			PostalCode:      "company_postal_code",
			Address:         "company_address",
		}
		companyID, err := h.companyRepository.Create(context.Background(), &company)
		if err != nil {
			t.Fatalf("err should be nil, but got: %v", err)
		}
		if companyID == 0 {
			t.Fatalf("domain.CompanyID should not be 0")
		}

		client := domain.Client{
			CompanyID:       companyID,
			Name:            "client_name",
			Representative:  "client_representative",
			TelephoneNumber: "client_telephone_number",
			PostalCode:      "client_postal_code",
			Address:         "client_address",
		}
		clientID, err := h.sub.Create(context.Background(), &client)
		if err != nil {
			t.Fatalf("err should be nil, but got: %v", err)
		}
		if clientID == 0 {
			t.Errorf("domain.ClientID should not be 0")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		h := createClientRepositoryTestHelper(t)

		client := domain.Client{
			CompanyID:       domain.CompanyID(10),
			Name:            "client_name",
			Representative:  "client_representative",
			TelephoneNumber: "client_telephone_number",
			PostalCode:      "client_postal_code",
			Address:         "client_address",
		}
		_, err := h.sub.Create(context.Background(), &client)
		if err == nil {
			t.Fatal("err should be occurred")
		}
	})
}
