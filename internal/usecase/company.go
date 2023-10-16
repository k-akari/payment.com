package usecase

import (
	"context"
	"fmt"

	"github.com/k-akari/payment.com/internal/domain"
)

type companyUsecase struct {
	companyRepository companyRepository
}

func NewCompanyUsecase(
	companyRepository companyRepository,
) *companyUsecase {
	return &companyUsecase{
		companyRepository: companyRepository,
	}
}

func (u *companyUsecase) CreateCompany(
	ctx context.Context,
	company *domain.Company,
) (domain.CompanyID, error) {
	cid, err := u.companyRepository.Create(ctx, company)
	if err != nil {
		return domain.CompanyID(0), fmt.Errorf("failed to run u.companyRepository.Create: %w", err)
	}

	return cid, nil
}

func (u *companyUsecase) GetCompanyByID(
	ctx context.Context,
	companyID domain.CompanyID,
) (*domain.Company, error) {
	company, err := u.companyRepository.GetByID(ctx, companyID)
	if err != nil {
		return nil, fmt.Errorf("failed to run u.companyRepository.GetByID: %w", err)
	}

	return company, nil
}
