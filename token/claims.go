package token

import (
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

// claims 相关，实现 IClaims
type (
	Claims struct {
		jwt.Claims
	}
	CFn func(claims *Claims)
)

func WithIssuer(issuer string) CFn {
	return func(claims *Claims) {
		claims.Issuer = issuer
	}
}
func WithClaimsID(Id string) CFn {
	return func(claims *Claims) {
		claims.ID = Id
	}
}
func WithSubject(subject string) CFn {
	return func(claims *Claims) {
		claims.Subject = subject
	}
}
func WithAudience(audience jwt.Audience) CFn {
	return func(claims *Claims) {
		claims.Audience = audience
	}
}
func WithExpiry(expiry *jwt.NumericDate) CFn {
	return func(claims *Claims) {
		claims.Expiry = expiry
	}
}
func WithNotBefore(notBefore *jwt.NumericDate) CFn {
	return func(claims *Claims) {
		claims.NotBefore = notBefore
	}
}
func WithIssuedAt(issuedAt *jwt.NumericDate) CFn {
	return func(claims *Claims) {
		claims.IssuedAt = issuedAt
	}
}

func NewClaims(fns ...CFn) *Claims {
	claims := &Claims{Claims: jwt.Claims{
		Expiry:    jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
	}}

	for _, fn := range fns {
		fn(claims)
	}

	return claims
}
func (c *Claims) Validated(v jwt.Expected) error {
	return c.Validate(v)
}
func (c *Claims) WithIssuer(issuer string) IClaims {
	c.Issuer = issuer
	return c
}
func (c *Claims) WithClaimsID(Id string) IClaims {
	c.ID = Id
	return c
}
func (c *Claims) WithSubject(subject string) IClaims {
	c.Subject = subject
	return c
}
func (c *Claims) WithAudience(audience jwt.Audience) IClaims {
	c.Audience = audience
	return c
}
func (c *Claims) WithExpiry(expiry *jwt.NumericDate) IClaims {
	c.Expiry = expiry
	return c
}
func (c *Claims) WithNotBefore(notBefore *jwt.NumericDate) IClaims {
	c.NotBefore = notBefore
	return c
}
func (c *Claims) WithIssuedAt(issuedAt *jwt.NumericDate) IClaims {
	c.IssuedAt = issuedAt
	return c
}

func (c *Claims) GetIssuer() string {
	return c.Issuer
}
func (c *Claims) GetID() string {
	return c.ID
}
func (c *Claims) GetSubject() string {
	return c.Subject
}
func (c *Claims) GetAudience() jwt.Audience {
	return c.Audience
}
func (c *Claims) GetExpiry() *jwt.NumericDate {
	return c.Expiry
}
func (c *Claims) GetNotBefore() *jwt.NumericDate {
	return c.NotBefore
}
func (c *Claims) GetIssuedAt() *jwt.NumericDate {
	return c.IssuedAt
}
