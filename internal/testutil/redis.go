package testutil

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	e := loadEnv()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", e.RedisHost, e.RedisPort),
		Password: "",
		DB:       0, // default database number
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect redis: %s", err)
	}

	return client
}
