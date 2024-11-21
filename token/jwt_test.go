package token

import (
	"github.com/google/uuid"
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
	)

	claims := NewClaims(
		WithIssuer("test-issuer"),
		WithSubject("test-subject"),
	)
	token := NewEngine[*Key, *User, *Claims](
		WithKey[*Key, *User, *Claims](keys),
		WithUser[*Key, *User, *Claims](user),
		WithClaims[*Key, *User, *Claims](claims),
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
	t.Logf("verifier: %+v %+v %+v %+v %+v", verifier.User().ID(), verifier.User().OpenID(), verifier.User().Email(), verifier.User().IsRefresh(), verifier.Claims().GetExpiry().Time())

	verifierR, err := token.VerifierRefresh(te)
	if err != nil {
		t.Logf("te VerifierRefresh err: %+v", err)
	}

	t.Logf("refresh: %s", re)
	verifierR, err = token.VerifierRefresh(re)
	if err != nil {
		panic(err)
	}
	t.Logf("verifierR: %+v %+v %+v %+v %+v", verifierR.User().ID(), verifierR.User().OpenID(), verifierR.User().Email(), verifierR.User().IsRefresh(), verifierR.Claims().GetExpiry().Time())

}
