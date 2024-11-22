package token

import (
	"crypto/rsa"
	"gopkg.in/square/go-jose.v2/jwt"
)

type (
	IKey interface {
		Private() *rsa.PrivateKey
		Public() *rsa.PublicKey
	}
	IUser interface {
		ID() uint64
		Email() string
		OpenID() string
		IsRefresh() bool
		WithIsRefresh(r bool) IUser
	}
	IClaims interface {
		GetIssuer() string
		GetID() string
		GetSubject() string
		GetAudience() jwt.Audience
		GetExpiry() *jwt.NumericDate
		GetNotBefore() *jwt.NumericDate
		GetIssuedAt() *jwt.NumericDate

		WithIssuer(issuer string) IClaims
		WithClaimsID(Id string) IClaims
		WithSubject(subject string) IClaims
		WithAudience(audience jwt.Audience) IClaims
		WithExpiry(expiry *jwt.NumericDate) IClaims
		WithNotBefore(notBefore *jwt.NumericDate) IClaims
		WithIssuedAt(issuedAt *jwt.NumericDate) IClaims
		Validated(v jwt.Expected) error
	}
	IEngine[K IKey, U IUser, C IClaims] interface {
		Key() K
		User() U
		Claims() C
		Generate() (token string, err error)
		VerifierToken(token string) (*U, *C, error)
		VerifierRefresh(token string) (*U, *C, error)
		WithKey(key K) IEngine[K, U, C]
		WithUser(user U) IEngine[K, U, C]
		WithClaims(claims C) IEngine[K, U, C]
	}
)
