package usecase

import (
	"context"
	"fmt"

	"github.com/k-akari/payment.com/internal/domain"
)

type clientUsecase struct {
	clientRepository clientRepository
}

func NewClientUsecase(
	clientRepository clientRepository,
) *clientUsecase {
	return &clientUsecase{
		clientRepository: clientRepository,
	}
}

func (u *clientUsecase) Create(
	ctx context.Context,
	client *domain.Client,
) (domain.ClientID, error) {
	cid, err := u.clientRepository.Create(ctx, client)
	if err != nil {
		return domain.ClientID(0), fmt.Errorf("failed to run u.clientRepository.Create: %w", err)
	}

	return cid, nil
}

func (u *clientUsecase) GetByID(
	ctx context.Context,
	clientID domain.ClientID,
) (*domain.Client, error) {
	client, err := u.clientRepository.GetByID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to run u.clientRepository.GetByID: %w", err)
	}

	return client, nil
}
