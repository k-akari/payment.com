//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock
package handler

import (
	"context"

	"github.com/k-akari/payment.com/internal/domain"
)

type (
	companyUsecase interface {
		CreateCompany(ctx context.Context, company *domain.Company) error
	}
)
