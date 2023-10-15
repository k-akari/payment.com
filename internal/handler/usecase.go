//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock
package handler

import (
	"context"

	"github.com/k-akari/payment.com/internal/domain"
)

type (
	companyUsecase interface {
		CreateCompany(ctx context.Context, company *domain.Company) (domain.CompanyID, error)
		GetCompanyByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error)
	}
	clientUsecase interface {
		Create(ctx context.Context, client *domain.Client) (domain.ClientID, error)
		GetByID(ctx context.Context, clientID domain.ClientID) (*domain.Client, error)
	}
)
