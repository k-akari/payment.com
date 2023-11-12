package auth

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/infrastructure/kvs"
	"github.com/k-akari/golang-rest-api-sample/internal/pkg/clock"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil/fixture"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
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
	sut, err := NewJWTer(&kvs.KVS{Cli: cli}, &clock.FixedClocker{})
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

func TestJWTer_GetToken(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cli := testutil.OpenRedisForTest(t)
	sut, err := NewJWTer(&kvs.KVS{Cli: cli}, &clock.FixedClocker{})
	if err != nil {
		t.Fatalf("want no error, but got %v", err)
	}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		// create token
		u := fixture.User(&domain.User{ID: domain.UserID(1234)})
		tok, err := jwt.NewBuilder().
			JwtID(uuid.New().String()).
			Issuer("golang-rest-api-sample").
			Subject("access_token").
			IssuedAt(sut.Clocker.Now()).
			Expiration(sut.Clocker.Now().Add(1*time.Hour)).
			Claim(UserNameKey, u.Name).
			Build()
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}

		// save token
		if err := sut.SessionStore.SaveUserID(ctx, tok.JwtID(), u.ID); err != nil {
			t.Fatalf("want no error, but got %v", err)
		}

		// sign token
		signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, sut.PrivateKey))
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}

		// set token to request header
		req := httptest.NewRequest(http.MethodGet, "https://github.com/k-akari", http.NoBody)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", string(signed)))

		got, err := sut.GetToken(ctx, req)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}
		if !reflect.DeepEqual(got, tok) {
			t.Errorf("want %v, but got %v", tok, got)
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		t.Run("unexisted token", func(t *testing.T) {
			t.Parallel()

			// create token
			u := fixture.User(&domain.User{ID: domain.UserID(12345)})
			tok, err := jwt.NewBuilder().
				JwtID(uuid.New().String()).
				Issuer("golang-rest-api-sample").
				Subject("access_token").
				IssuedAt(sut.Clocker.Now()).
				Expiration(sut.Clocker.Now().Add(1*time.Hour)).
				Claim(UserNameKey, u.Name).
				Build()
			if err != nil {
				t.Fatalf("want no error, but got %v", err)
			}

			// sign token
			signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, sut.PrivateKey))
			if err != nil {
				t.Fatalf("want no error, but got %v", err)
			}

			// set token to request header
			req := httptest.NewRequest(http.MethodGet, "https://github.com/k-akari", http.NoBody)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", string(signed)))

			_, err = sut.GetToken(ctx, req)
			if !errors.Is(err, redis.Nil) {
				t.Errorf("want %v, but got %v", redis.Nil, err)
			}
		})
	})
}
