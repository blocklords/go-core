package token

import (
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/jwt"
	"testing"
)

func TestNewEngine(t *testing.T) {
	openId := uuid.New()
	t.Logf("openId: %+v", openId.String())

	keys := NewKey()
	user := NewUser(
		WithID(333),
		WithOpenId(openId),
		WithEmail(`123@gmail.com`),
		WithClaims(jwt.Claims{
			Issuer:  "test-1",
			Subject: "test-1",
		}),
	)
	token := NewEngine(
		WithKey(keys),
		WithUser(
			user,
		),
	)

	te, re, err := token.Generate()
	if err != nil {
		panic(err)
	}

	t.Logf("token: %s", te)

	verifier, err := token.VerifierToken(te)
	if err != nil {
		panic(err)
	}
	t.Logf("verifier: %+v %+v %+v %+v %+v", verifier.ID(), verifier.OpenID(), verifier.Email(), verifier.IsRefresh(), verifier.Claims().Expiry.Time())

	verifierR, err := token.VerifierRefresh(te)
	if err != nil {
		t.Logf("te VerifierRefresh err: %+v", err)
	}

	t.Logf("refresh: %s", re)
	verifierR, err = token.VerifierRefresh(re)
	if err != nil {
		panic(err)
	}
	t.Logf("verifierR: %+v %+v %+v %+v %+v", verifierR.ID(), verifierR.OpenID(), verifierR.Email(), verifierR.IsRefresh(), verifierR.Claims().Expiry.Time())

}
