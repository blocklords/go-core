package token

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"log"
	"strings"
	"time"
)

type (
	IClaims interface {
		ID() uint64
		Email() string
		OpenID() string
	}

	IJWT interface {
		PrivateKey() *rsa.PrivateKey
		PublicKey() *rsa.PublicKey
		AccessMaxAge() time.Duration
		RefreshMaxAge() time.Duration
	}
)

const (
	priString          = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBcDg1S1Vvd1BvSEh5WmxnNC9xUG5JdmJCOUdyLy9zTWpTVEhHeGtOeUZpS2Z3Z0ZFCkJyTkFRUjF2ekMxQzVyTUM5S3pITFRPZXhKSHJxL0FrVU84eWhyb1ZzaWJOQllFRkxTdExqMkMzdVhrUUwxYlgKR0VLai9MWGpBVStzYXYxMEFXL1VLRnpJT2ZONFFUbVBLVWgxQTdzK3EyRUVvS3BtZUZrZWt5WU03bkJNemxoMAo4bmxFWm\n1LYU5TcHBMOGZxM0pOcEtpS1AxWjZvWnhYeHBuTzdRTUZPTVZOZzVWRFZodjBweE4wRzA5RjN6OEtvCldQeEY1NXFKdS9QU2RlbjFMc1FsS3ZFaDMzZXY0QmdTRnV1UHRKR0xCaDlGazE3Q3Z1V2loTkdrd0dGY2MyQTUKaW9DOVVZamhKWlM3UXhwOEVWdHViRkV6bGk1Zk1HL211eXpJRFFJREFRQUJBb0lCQVFDUjVWWFJOMTFPemtORwpvR1hOWDR2U1ptQnp0YVFsU0Z3aGcxbWpmM2h0cm1UZ05HR0V3Y3lYMEpR\nbkhTTVJtWXAwV05SRGhLSUJuaTBkCkxJa21wUkYwK2Mxck96aitGQk1BRnFoM1hFdmd3bFZFRTJpbit5akF5eE0zVEtKSDAxMU04b0d2Smh3ZjVvTWoKa252YUZObElDVUNQbUthQldpWUZkTmFVY1h6RXdRNUEyQXBOeWM4RlJzNlh2RVY1dFhueXNMR0lub0RoeGZCVQpOeUIvYStCWlplRkpudEc0Q1h0Q1Ewb3BoQldWSjk4NHAvRTdZTitPNE5lOUdJdDJ6YW9oTjJJbXFHK1psUkFzCkZqb1M5aWxqbnhSNmg5c1\npiYnBERzVlcHNFeXJtUWlHQUpGTURzV3YvNTUyZzZtMFNUaGNjcVpFeHhUNDg0N1YKYVQzVlFNVUJBb0dCQU5SK29pdHhJcU1ieTFjb2Z3U2xtcVR0RURLZmNnc2tPaXBjb2VhOUw1NUp0WXNKZExzSAp0eUphd3l5U1IvalVXeUZnaEZodUFzR2ZiakJMRWQ4NTZZQXBRcFlQbDBwcnpNaklWVStISmdhTnZGRFdXa2h4CnJKb3p0V1crQ1pPb293VXdMTXljSEJKVWpaVVBYWEhJNmpyQ2xVU0xibDVHcDUvZWN4NVVR\nV1J4QW9HQkFNb3AKWjdQc0xVRzA0ZVdmcVVSRnc5K2Y3SlBKMjdKSTRKTDhUL2lmZmhiZDkxRW1DVUs1T3ZEWjQ1QkZPZTA3RlkrWgpQdk5lZ0k2bkJnWGlVTS9qVHdxQzJUbTNac0MxakE1a0JSVUNHbnZ2VVZZSTVXLy9kUUR6Q1NmUStLc1dyUEcyCkk4L0p2K1pqZTh4ZnJEbUpmS0Uvb1BnTC9qSVBXVWY0RTd4SE9YdGRBb0dCQUltSUFJd2ZhR3lyUTV1QXdWMFAKbGh5aXRyWWREcUg1YTVBWmJrdzZMRVRGcm\npOMEJsSTY5eVBNSE1DUFdQZks4Y1NUaEhUN2x0c2N4aU9Kb3VLWQp4L0ZFUXkxK244dnlJNVBjWGFMZ2RSTU96MUIrdStaaGRIWkZlMldEYncxYnUwOVRVOXVHT29EK3FyaE1QbzJ6Cm5TNDAzSW1GdVFSWnRJbzdYc1RGZ2FGeEFvR0JBTWxMZ1I3K1EvSHhFaDE2WlNpOTd0TjBnalNHQW1QN2ZPSGUKcWlKOWJTZUh6UUxZUk5CVGNBVHljRXp2SVVhK1ZqR3QvYWlHcUt0aVUvVDM3RStUbnRod2dhdWVtb200TzhU\nNApncmJ3YVQ2T2hRYU54U2RIemxFcnJpb2ZRZnZaa0VyOWVac2s0QmVmWjEyUXhnUmtpZHhsWnZxVnRuNVNHaXczCk1DK0JIQk5oQW9HQUZyVEZDTllJemROcDB4eE1xc05ISGZaL1FGSGszb2w3eXdEVU4wOXZNeFhyVkZ0UE8xTVQKNEtydGEwQ3BrY2RINmgrZU9XbGdoWTlCczhpbVIySlZxSUJBTnRBd0kza2RiNG5aQm5NOTkzS0lzMFlGcWg0LwpOdHZHL3g0ME5CU2JQQklDZWE4RnFBRGpmaEhHTkRYa080NW\ngxTFovOGJFSU9VK1BKNVc3RmxNPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
	pubString          = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFwODVLVW93UG9ISHlabGc0L3FQbgpJdmJCOUdyLy9zTWpTVEhHeGtOeUZpS2Z3Z0ZFQnJOQVFSMXZ6QzFDNXJNQzlLekhMVE9leEpIcnEvQWtVTzh5Cmhyb1ZzaWJOQllFRkxTdExqMkMzdVhrUUwxYlhHRUtqL0xYakFVK3NhdjEwQVcvVUtGeklPZk40UVRtUEtVaDEKQTdzK3EyRUVvS3\nBtZUZrZWt5WU03bkJNemxoMDhubEVabUthTlNwcEw4ZnEzSk5wS2lLUDFaNm9aeFh4cG5PNwpRTUZPTVZOZzVWRFZodjBweE4wRzA5RjN6OEtvV1B4RjU1cUp1L1BTZGVuMUxzUWxLdkVoMzNldjRCZ1NGdXVQCnRKR0xCaDlGazE3Q3Z1V2loTkdrd0dGY2MyQTVpb0M5VVlqaEpaUzdReHA4RVZ0dWJGRXpsaTVmTUcvbXV5ekkKRFFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	accessTokenMaxAge  = 24 * time.Hour
	refreshTokenMaxAge = 24 * 90 * time.Hour
)

type (
	DefaultUser struct {
		Id     uint64 `json:"id"`
		Mail   string `json:"email"`
		OpenId string `json:"openId"`
	}
	UFn           func(u *DefaultUser)
	DefaultEngine struct {
		private                     *rsa.PrivateKey
		public                      *rsa.PublicKey
		accessMaxAge, refreshMaxAge time.Duration
	}
	EFn func(e *DefaultEngine)
)

func WithPrivateKey(private *rsa.PrivateKey) EFn {
	return func(e *DefaultEngine) {
		e.private = private
	}
}
func WithPublicKey(public *rsa.PublicKey) EFn {
	return func(e *DefaultEngine) {
		e.public = public
	}
}
func WithAccessMax(access time.Duration) EFn {
	return func(e *DefaultEngine) {
		e.accessMaxAge = access
	}
}
func WithRefreshMax(refresh time.Duration) EFn {
	return func(e *DefaultEngine) {
		e.refreshMaxAge = refresh
	}
}

func NewEngine(fns ...EFn) *DefaultEngine {
	private, _ := base64.StdEncoding.DecodeString(priString)
	privateKey, _ := jwt.ParsePrivateKeyRSA(private)
	public, _ := base64.StdEncoding.DecodeString(pubString)
	publicKey, _ := jwt.ParsePublicKeyRSA(public)
	engine := &DefaultEngine{
		private:       privateKey,
		public:        publicKey,
		accessMaxAge:  accessTokenMaxAge,
		refreshMaxAge: refreshTokenMaxAge,
	}
	for _, fn := range fns {
		fn(engine)
	}
	return engine
}
func (d *DefaultEngine) PrivateKey() *rsa.PrivateKey {
	return d.private
}
func (d *DefaultEngine) PublicKey() *rsa.PublicKey {
	return d.public
}
func (d *DefaultEngine) AccessMaxAge() time.Duration {
	return d.accessMaxAge
}
func (d *DefaultEngine) RefreshMaxAge() time.Duration {
	return d.refreshMaxAge
}

func WithUserID(ID uint64) UFn {
	return func(u *DefaultUser) {
		u.Id = ID
	}
}
func WithEmail(email string) UFn {
	return func(u *DefaultUser) {
		u.Mail = email
	}
}
func WithOpenId(openId string) UFn {
	return func(u *DefaultUser) {
		u.OpenId = openId
	}
}
func NewUser(fns ...UFn) *DefaultUser {
	user := &DefaultUser{}
	for _, fn := range fns {
		fn(user)
	}
	return user
}
func (u *DefaultUser) ID() uint64 {
	return u.Id
}
func (u *DefaultUser) Email() string {
	return u.Mail
}
func (u *DefaultUser) OpenID() string {
	return u.OpenId
}

type JWT[M IClaims, E IJWT] struct {
	model    M
	engine   E
	signer   *jwt.Signer
	verifier *jwt.Verifier
}

func NewJWT[M IClaims, E IJWT](model M, engine E) *JWT[M, E] {
	return &JWT[M, E]{
		model:    model,
		engine:   engine,
		signer:   jwt.NewSigner(jwt.RS256, engine.PrivateKey(), engine.AccessMaxAge()),
		verifier: jwt.NewVerifier(jwt.RS256, engine.PublicKey()),
	}
}

func (j *JWT[M, E]) Generate() (access, refresh string, err error) {
	now := time.Now()
	tokenPair, err := j.signer.NewTokenPair(j.model, struct {
		claims    M
		isRefresh bool
	}{
		claims:    j.model,
		isRefresh: true,
	}, j.engine.RefreshMaxAge(), jwt.Claims{NotBefore: now.Unix(), IssuedAt: now.Unix(), Expiry: now.Add(j.engine.AccessMaxAge()).Unix()})
	if err != nil {
		return "", "", err
	}
	return strings.ReplaceAll(string(tokenPair.AccessToken), `"`, ""), strings.ReplaceAll(string(tokenPair.RefreshToken), `"`, ""), nil
}

func (j *JWT[M, E]) Verify(ctx iris.Context) {
	authorization := strings.Split(ctx.GetHeader("Authorization"), " ")
	token := strings.ReplaceAll(authorization[len(authorization)-1], "Bearer", "")

	verifiedToken, err := j.verifier.VerifyToken([]byte(token))
	if err != nil {
		ctx.StopWithProblem(iris.StatusUnauthorized, iris.NewProblem().Title(`Unauthorized`))
		return
	}

	var claims M
	if err = verifiedToken.Claims(&claims); err != nil {
		ctx.StopWithProblem(iris.StatusForbidden, iris.NewProblem().Title(`token is invalid`))
		return
	}

	ctx.Values().Set("claims", claims)

	ctx.Next()
}

func Verify[M IClaims, E IJWT](j *JWT[M, E], token string) {
	verifiedToken, err := j.verifier.VerifyToken([]byte(token))
	if err != nil {
		panic(err)
	}

	var claims M
	if err = verifiedToken.Claims(&claims); err != nil {
		panic(err)
	}

	log.Printf("claims: %+v", claims)
}
