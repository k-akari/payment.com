//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock
package usecase

import (
	"context"

	"github.com/k-akari/payment.com/internal/domain"
)

type (
	companyRepository interface {
		Create(ctx context.Context, company *domain.Company) (domain.CompanyID, error)
		GetByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error)
	}
	clientRepository interface {
		Create(ctx context.Context, client *domain.Client) (domain.ClientID, error)
		GetByID(ctx context.Context, clientID domain.ClientID) (*domain.Client, error)
	}
)
