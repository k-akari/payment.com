package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/usecase/mock"
	gomock "go.uber.org/mock/gomock"
)

type companyUsecaseTestHelper struct {
	sub               *adminUsecase
	ctx               context.Context
	companyRepository *mock.MockcompanyRepository
}

func createCompanyUsecaseTestHelper(t *testing.T) *companyUsecaseTestHelper {
	t.Helper()

	ctrl := gomock.NewController(t)

	companyRepository := mock.NewMockcompanyRepository(ctrl)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return &companyUsecaseTestHelper{
		sub:               NewCompanyUsecase(companyRepository),
		ctx:               context.Background(),
		companyRepository: companyRepository,
	}
}

func TestAdminUsecase_CreateCompany(t *testing.T) {
	t.Parallel()

	company := domain.Company{
		Name:            "name",
		Representative:  "representative",
		TelephoneNumber: "telephone_number",
		PostalCode:      "postal_code",
		Address:         "address",
	}

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		h := createCompanyUsecaseTestHelper(t)

		cid := domain.CompanyID(10)
		h.companyRepository.EXPECT().Create(h.ctx, &company).Return(cid, nil)

		got, err := h.sub.CreateCompany(h.ctx, &company)
		if err != nil {
			t.Errorf("err should be nil, but got: %v", err)
		}
		if got != cid {
			t.Errorf("got: %v, want: %v", got, cid)
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		h := createCompanyUsecaseTestHelper(t)

		h.companyRepository.EXPECT().Create(h.ctx, &company).Return(domain.CompanyID(0), errors.New("error"))

		_, err := h.sub.CreateCompany(h.ctx, &company)
		if err == nil {
			t.Error("err should not be nil")
		}
	})
}
