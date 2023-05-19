package token

import (
	"log"
	"time"

	"github.com/cristalhq/jwt/v4"
)

type Manager struct {
	Secret        []byte
	SigningMethod jwt.Algorithm
	Verifier      jwt.Verifier
}

func NewManager(secret []byte, SigningMethod jwt.Algorithm) *Manager {
	return &Manager{
		Secret:        secret,
		SigningMethod: SigningMethod,
	}
}
func (tm *Manager) GenerateToken(userId string, userName string, TTL time.Duration) (*jwt.Token, error) {

	signer, err := jwt.NewSignerHS(tm.SigningMethod, tm.Secret)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// create claims (you can create your own, see: Example_BuildUserClaims)
	claims := &jwt.RegisteredClaims{
		ID:        userId,
		Subject:   userName,
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(TTL)},
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
	}

	// create a Builder
	builder := jwt.NewBuilder(signer)

	// and build a Token
	token, err := builder.Build(claims)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return token, nil
}
