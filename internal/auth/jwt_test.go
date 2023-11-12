package auth

import (
	"bytes"
	"context"
	"testing"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/infrastructure/kvs"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil/fixture"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPubKey, want) {
		t.Errorf("want %s, but got %s", want, rawPubKey)
	}
	want = []byte("-----BEGIN PRIVATE KEY-----")
	if !bytes.Contains(rawPrivKey, want) {
		t.Errorf("want %s, but got %s", want, rawPrivKey)
	}
}

func TestJWTer_CreateToken(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cli := testutil.OpenRedisForTest(t)
	sut, err := NewJWTer(&kvs.KVS{Cli: cli})
	if err != nil {
		t.Fatalf("want no error, but got %v", err)
	}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		uid := domain.UserID(1234)
		u := fixture.User(&domain.User{ID: uid})
		got, err := sut.CreateToken(ctx, u)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}
		if len(got) == 0 {
			t.Errorf("want len(got) > 0, but got %d", len(got))
		}
	})
}
