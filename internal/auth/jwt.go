package auth

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/pkg/clock"
	"github.com/lestrrat-go/jwx/v2/jwk"
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
