package auth

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/pkg/clock"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const (
	UserNameKey = "user_name"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	SessionStore          SessionStore
	Clocker               clock.Clocker
}

type SessionStore interface {
	SaveUserID(ctx context.Context, key string, userID domain.UserID) error
	LoadUserID(ctx context.Context, key string) (domain.UserID, error)
}

func NewJWTer(s SessionStore) (*JWTer, error) {
	j := &JWTer{SessionStore: s, Clocker: clock.RealClocker{}}

	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	j.PrivateKey = privkey

	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}
	j.PublicKey = pubkey

	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (j *JWTer) CreateToken(ctx context.Context, u *domain.User) ([]byte, error) {
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer("golang-rest-api-sample").
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(1*time.Hour)).
		Claim(UserNameKey, u.Name).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed in CreateToken: %w", err)
	}

	if err := j.SessionStore.SaveUserID(ctx, tok.JwtID(), u.ID); err != nil {
		return nil, fmt.Errorf("failed in SaveUserID: %w", err)
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, fmt.Errorf("failed in Sign: %w", err)
	}

	return signed, nil
}
