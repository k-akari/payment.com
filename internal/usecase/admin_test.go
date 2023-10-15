package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/usecase/mock"
	gomock "go.uber.org/mock/gomock"
)

type adminUsecaseTestHelper struct {
	sub               *adminUsecase
	ctx               context.Context
	companyRepository *mock.MockcompanyRepository
}

func createAdminUsecaseTestHelper(t *testing.T) *adminUsecaseTestHelper {
	t.Helper()

	ctrl := gomock.NewController(t)

	companyRepository := mock.NewMockcompanyRepository(ctrl)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return &adminUsecaseTestHelper{
		sub:               NewAdminUsecase(companyRepository),
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

		h := createAdminUsecaseTestHelper(t)

		h.companyRepository.EXPECT().Create(h.ctx, &company).Return(nil)

		err := h.sub.CreateCompany(h.ctx, &company)
		if err != nil {
			t.Errorf("err should be nil, but got: %v", err)
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		h := createAdminUsecaseTestHelper(t)

		h.companyRepository.EXPECT().Create(h.ctx, &company).Return(errors.New("error"))

		err := h.sub.CreateCompany(h.ctx, &company)
		if err == nil {
			t.Error("err should not be nil")
		}
	})
}
