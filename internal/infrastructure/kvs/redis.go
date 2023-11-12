package kvs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/k-akari/golang-rest-api-sample/internal/domain"
)

type KVS struct {
	Cli *redis.Client
}

func NewClient(ctx context.Context, host string, port int) (*KVS, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &KVS{Cli: cli}, nil
}

func (k *KVS) SaveUserID(ctx context.Context, key string, userID domain.UserID) error {
	id := int64(userID)
	if err := k.Cli.Set(ctx, key, id, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to set %q: %w", key, err)
	}

	return nil
}

func (k *KVS) LoadUserID(ctx context.Context, key string) (domain.UserID, error) {
	id, err := k.Cli.Get(ctx, key).Int64()
	if err != nil {
		return 0, fmt.Errorf("failed to get by %q: %w", key, err)
	}

	return domain.UserID(id), nil
}
