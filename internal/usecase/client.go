package usecase

import (
	"context"
	"fmt"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
)

type ClientUsecase struct {
	clientRepository clientRepository
}

func NewClientUsecase(
	clientRepository clientRepository,
) *ClientUsecase {
	return &ClientUsecase{
		clientRepository: clientRepository,
	}
}

func (u *ClientUsecase) Create(
	ctx context.Context,
	client *domain.Client,
) (domain.ClientID, error) {
	cid, err := u.clientRepository.Create(ctx, client)
	if err != nil {
		return domain.ClientID(0), fmt.Errorf("failed to run u.clientRepository.Create: %w", err)
	}

	return cid, nil
}

func (u *ClientUsecase) GetByID(
	ctx context.Context,
	clientID domain.ClientID,
) (*domain.Client, error) {
	client, err := u.clientRepository.GetByID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to run u.clientRepository.GetByID: %w", err)
	}

	return client, nil
}
