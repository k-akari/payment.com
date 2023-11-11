package kvs

import (
	"context"
	"testing"
	"time"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil"
)

func TestKVS_Save(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)
	sut := &KVS{cli: cli}

	key := "TestKVS_Save"
	uid := domain.UserID(1234)
	ctx := context.Background()

	t.Cleanup(func() { cli.Del(ctx, key) })

	if err := sut.SaveUserID(ctx, key, uid); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

func TestKVS_Load(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)
	sut := &KVS{cli: cli}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		key := "TestKVS_Load_ok"
		uid := domain.UserID(1234)
		ctx := context.Background()
		cli.Set(ctx, key, int64(uid), 30*time.Minute)
		t.Cleanup(func() { cli.Del(ctx, key) })
		got, err := sut.LoadUserID(ctx, key)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}
		if got != uid {
			t.Errorf("want %d, but got %d", uid, got)
		}
	})

	t.Run("notFound", func(t *testing.T) {
		t.Parallel()

		key := "TestKVS_Save_notFound"
		ctx := context.Background()
		_, err := sut.LoadUserID(ctx, key)
		if err == nil {
			t.Errorf("want error, but got nil")
		}
	})
}
