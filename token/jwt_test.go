package token

import (
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/jwt"
	"testing"
	"time"
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

	te, err := token.Generate()
	if err != nil {
		panic(err)
	}

	t.Logf("token: %s", te)

	u, c, err := token.VerifierToken(te)
	if err != nil {
		panic(err)
	}

	us, cl := *u, *c
	t.Logf("verifier: %+v %+v %+v %+v %+v", us.ID(), us.OpenID(), us.Email(), us.IsRefresh(), cl.GetExpiry().Time())

	ru, rc, err := token.VerifierRefresh(te)
	if err != nil {
		t.Logf("te VerifierRefresh err: %+v", err)
	}

	claims.WithExpiry(jwt.NewNumericDate(time.Now().UTC().Add(7 * 24 * time.Hour)))
	user.WithIsRefresh(true)
	re, err := token.Generate()
	if err != nil {
		panic(err)
	}

	t.Logf("refresh: %s", re)
	ru, rc, err = token.VerifierRefresh(re)
	if err != nil {
		panic(err)
	}

	rus, rcl := *ru, *rc
	t.Logf("verifierR: %+v %+v %+v %+v %+v", rus.ID(), rus.OpenID(), rus.Email(), rus.IsRefresh(), rcl.GetExpiry().Time())

}
