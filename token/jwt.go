package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/blocklords/go-core/entity"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

const (
	private = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAp85KUowPoHHyZlg4/qPnIvbB9Gr//sMjSTHGxkNyFiKfwgFE\nBrNAQR1vzC1C5rMC9KzHLTOexJHrq/AkUO8yhroVsibNBYEFLStLj2C3uXkQL1bX\nGEKj/LXjAU+sav10AW/UKFzIOfN4QTmPKUh1A7s+q2EEoKpmeFkekyYM7nBMzlh0\n8nlEZmKaNSppL8fq3JNpKiKP1Z6oZxXxpnO7QMFOMVNg5VDVhv0pxN0G09F3z8Ko\nWPxF55qJu/PSden1LsQlKvEh33ev4BgSFuuPtJGLBh9Fk17CvuWihNGkwGFcc2A5\nioC9UYjhJZS7Qxp8EVtubFEzli5fMG/muyzIDQIDAQABAoIBAQCR5VXRN11OzkNG\noGXNX4vSZmBztaQlSFwhg1mjf3htrmTgNGGEwcyX0JQnHSMRmYp0WNRDhKIBni0d\nLIkmpRF0+c1rOzj+FBMAFqh3XEvgwlVEE2in+yjAyxM3TKJH011M8oGvJhwf5oMj\nknvaFNlICUCPmKaBWiYFdNaUcXzEwQ5A2ApNyc8FRs6XvEV5tXnysLGInoDhxfBU\nNyB/a+BZZeFJntG4CXtCQ0ophBWVJ984p/E7YN+O4Ne9GIt2zaohN2ImqG+ZlRAs\nFjoS9iljnxR6h9sZbbpDG5epsEyrmQiGAJFMDsWv/552g6m0SThccqZExxT4847V\naT3VQMUBAoGBANR+oitxIqMby1cofwSlmqTtEDKfcgskOipcoea9L55JtYsJdLsH\ntyJawyySR/jUWyFghFhuAsGfbjBLEd856YApQpYPl0przMjIVU+HJgaNvFDWWkhx\nrJoztWW+CZOoowUwLMycHBJUjZUPXXHI6jrClUSLbl5Gp5/ecx5UQWRxAoGBAMop\nZ7PsLUG04eWfqURFw9+f7JPJ27JI4JL8T/iffhbd91EmCUK5OvDZ45BFOe07FY+Z\nPvNegI6nBgXiUM/jTwqC2Tm3ZsC1jA5kBRUCGnvvUVYI5W//dQDzCSfQ+KsWrPG2\nI8/Jv+Zje8xfrDmJfKE/oPgL/jIPWUf4E7xHOXtdAoGBAImIAIwfaGyrQ5uAwV0P\nlhyitrYdDqH5a5AZbkw6LETFrjN0BlI69yPMHMCPWPfK8cSThHT7ltscxiOJouKY\nx/FEQy1+n8vyI5PcXaLgdRMOz1B+u+ZhdHZFe2WDbw1bu09TU9uGOoD+qrhMPo2z\nnS403ImFuQRZtIo7XsTFgaFxAoGBAMlLgR7+Q/HxEh16ZSi97tN0gjSGAmP7fOHe\nqiJ9bSeHzQLYRNBTcATycEzvIUa+VjGt/aiGqKtiU/T37E+Tnthwgauemom4O8T4\ngrbwaT6OhQaNxSdHzlErriofQfvZkEr9eZsk4BefZ12QxgRkidxlZvqVtn5SGiw3\nMC+BHBNhAoGAFrTFCNYIzdNp0xxMqsNHHfZ/QFHk3ol7ywDUN09vMxXrVFtPO1MT\n4Krta0CpkcdH6h+eOWlghY9Bs8imR2JVqIBANtAwI3kdb4nZBnM993KIs0YFqh4/\nNtvG/x40NBSbPBICea8FqADjfhHGNDXkO45h1LZ/8bEIOU+PJ5W7FlM=\n-----END RSA PRIVATE KEY-----"
	public  = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp85KUowPoHHyZlg4/qPn\nIvbB9Gr//sMjSTHGxkNyFiKfwgFEBrNAQR1vzC1C5rMC9KzHLTOexJHrq/AkUO8y\nhroVsibNBYEFLStLj2C3uXkQL1bXGEKj/LXjAU+sav10AW/UKFzIOfN4QTmPKUh1\nA7s+q2EEoKpmeFkekyYM7nBMzlh08nlEZmKaNSppL8fq3JNpKiKP1Z6oZxXxpnO7\nQMFOMVNg5VDVhv0pxN0G09F3z8KoWPxF55qJu/PSden1LsQlKvEh33ev4BgSFuuP\ntJGLBh9Fk17CvuWihNGkwGFcc2A5ioC9UYjhJZS7Qxp8EVtubFEzli5fMG/muyzI\nDQIDAQAB\n-----END PUBLIC KEY-----"
)

type (
	IKey interface {
		Private() *rsa.PrivateKey
		Public() *rsa.PublicKey
	}
	Key struct {
		private *rsa.PrivateKey
		public  *rsa.PublicKey
	}

	KFn func(k *Key)
)

func WithPrivate(private *rsa.PrivateKey) KFn {
	return func(k *Key) {
		k.private = private
	}
}

func WithPublic(public *rsa.PublicKey) KFn {
	return func(k *Key) {
		k.public = public
	}
}

func (k *Key) Private() *rsa.PrivateKey {
	return k.private
}

func (k *Key) Public() *rsa.PublicKey {
	return k.public
}

func NewKey() *Key {
	privateBlock, _ := pem.Decode([]byte(private))
	if privateBlock == nil {
		panic(fmt.Errorf("private key: malformed or missing PEM format (RSA)"))
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		if keyP, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes); err == nil {
			pKey, ok := keyP.(*rsa.PrivateKey)
			if !ok {
				panic(fmt.Errorf("private key: expected a type of *rsa.PrivateKey"))
			}

			privateKey = pKey
		} else {
			panic(err)
		}
	}

	publicBlock, _ := pem.Decode([]byte(public))
	if privateBlock == nil {
		panic("jwt: Could not decode public key")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		if cert, err := x509.ParseCertificate(publicBlock.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			panic(err)
		}
	}

	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("public key: expected a type of *rsa.PublicKey"))
	}

	return &Key{
		private: privateKey,
		public:  publicKey,
	}
}

type (
	IUser interface {
		ID() uint64
		Email() string
		OpenID() string
		IsRefresh() bool
		WithIsRefresh(r bool) IUser
	}
	User struct {
		id             uint64             `json:"id"`
		salt           string             `json:"salt"`
		environment    entity.Environment `json:"environment"`
		isRefresh      bool               `json:"isRefresh"`
		email          string             `json:"email"`
		privyWallet    string             `json:"privyWallet"`
		changePassword bool               `json:"changePassword"`
		openId         uuid.UUID          `json:"openId"`
	}

	UFn func(u *User)
)

func WithID(id uint64) UFn {
	return func(u *User) {
		u.id = id
	}
}
func WithEmail(email string) UFn {
	return func(u *User) {
		u.email = email
	}
}
func WithOpenId(openId uuid.UUID) UFn {
	return func(u *User) {
		u.openId = openId
	}
}
func WithEnvironment(env entity.Environment) UFn {
	return func(u *User) {
		u.environment = env
	}
}
func WithIsRefresh(isRefresh bool) UFn {
	return func(u *User) {
		u.isRefresh = isRefresh
	}
}
func WithSalt(salt string) UFn {
	return func(u *User) {
		u.salt = salt
	}
}
func WithWallet(wallet string) UFn {
	return func(u *User) {
		u.privyWallet = wallet
	}
}
func WithChangePassword(change bool) UFn {
	return func(u *User) {
		u.changePassword = change
	}
}
func NewUser(fns ...UFn) *User {
	auth := &User{isRefresh: false}
	for _, fn := range fns {
		fn(auth)
	}
	return auth
}

// 链式操作
func (a *User) WithID(id uint64) *User {
	a.id = id
	return a
}
func (a *User) WithEmail(email string) *User {
	a.email = email
	return a
}
func (a *User) WithOpenID(openId uuid.UUID) *User {
	a.openId = openId
	return a
}
func (a *User) WithEnvironment(env entity.Environment) *User {
	a.environment = env
	return a
}
func (a *User) WithIsRefresh(r bool) IUser {
	a.isRefresh = r
	return a
}
func (a *User) WithSalt(salt string) *User {
	a.salt = salt
	return a
}
func (a *User) WithPrivyWallet(wallet string) *User {
	a.privyWallet = wallet
	return a
}
func (a *User) WithChangePassword(cp bool) *User {
	a.changePassword = cp
	return a
}

func (a *User) ID() uint64 {
	return a.id
}
func (a *User) Email() string {
	return a.email
}
func (a *User) OpenID() string {
	return a.openId.String()
}
func (a *User) Environment() entity.Environment {
	return a.environment
}
func (a *User) IsRefresh() bool {
	return a.isRefresh
}
func (a *User) Salt() string {
	return a.salt
}
func (a *User) PrivyWallet() string {
	return a.privyWallet
}
func (a *User) ChangePassword() bool {
	return a.changePassword
}

func (a *User) UnmarshalJSON(data []byte) error {
	temp := new(struct {
		Id             uint64             `json:"id"`
		Salt           string             `json:"salt"`
		Environment    entity.Environment `json:"environment"`
		IsRefresh      bool               `json:"isRefresh"`
		Email          string             `json:"email"`
		PrivyWallet    string             `json:"privyWallet"`
		ChangePassword bool               `json:"changePassword"`
		OpenId         uuid.UUID          `json:"openId"`
	})
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &temp); err != nil {
		return err
	}

	a.id = temp.Id
	a.salt = temp.Salt
	a.environment = temp.Environment
	a.isRefresh = temp.IsRefresh
	a.email = temp.Email
	a.privyWallet = temp.PrivyWallet
	a.changePassword = temp.ChangePassword
	a.openId = temp.OpenId
	return nil
}
func (a *User) MarshalJSON() ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(struct {
		Id             uint64             `json:"id"`
		Salt           string             `json:"salt"`
		Environment    entity.Environment `json:"environment"`
		IsRefresh      bool               `json:"isRefresh"`
		Email          string             `json:"email"`
		PrivyWallet    string             `json:"privyWallet"`
		ChangePassword bool               `json:"changePassword"`
		OpenId         uuid.UUID          `json:"openId"`
	}{
		Id:             a.id,
		Salt:           a.salt,
		Environment:    a.environment,
		IsRefresh:      a.isRefresh,
		Email:          a.email,
		PrivyWallet:    a.privyWallet,
		ChangePassword: a.changePassword,
		OpenId:         a.openId,
	})
}

type (
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
		Expiry:    jwt.NewNumericDate(time.Now().UTC().Add(2 * time.Hour)),
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

type (
	IEngine[K IKey, U IUser, C IClaims] interface {
		Key() K
		User() U
		Claims() C
		Generate() (token string, err error)
		VerifierToken(token string) (UserClaims[U, C], error)
		VerifierRefresh(token string) (UserClaims[U, C], error)
		WithKey(key K) IEngine[K, U, C]
		WithUser(user U) IEngine[K, U, C]
	}

	Engine[K IKey, U IUser, C IClaims] struct {
		key    K
		user   U
		claims C
	}
	EFn[K IKey, U IUser, C IClaims] func(e *Engine[K, U, C])

	UserClaims[U IUser, C IClaims] struct {
		user   U `json:"user"`
		claims C `json:"claims"`
	}
)

func (uc *UserClaims[U, C]) User() U {
	return uc.user
}

func (uc *UserClaims[U, C]) Claims() C {
	return uc.claims
}

func (uc *UserClaims[U, C]) UnmarshalJSON(data []byte) error {
	temp := new(struct {
		User   U `json:"user"`
		Claims C `json:"claims"`
	})
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &temp); err != nil {
		return err
	}

	uc.user = temp.User
	uc.claims = temp.Claims
	return nil
}
func (uc *UserClaims[U, C]) MarshalJSON() ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(struct {
		User   U `json:"user"`
		Claims C `json:"claims"`
	}{
		User:   uc.user,
		Claims: uc.claims,
	})
}

func WithKey[K IKey, U IUser, C IClaims](key K) EFn[K, U, C] {
	return func(e *Engine[K, U, C]) {
		e.key = key
	}
}
func WithUser[K IKey, U IUser, C IClaims](user U) EFn[K, U, C] {
	return func(e *Engine[K, U, C]) {
		e.user = user
	}
}
func WithClaims[K IKey, U IUser, C IClaims](claims C) EFn[K, U, C] {
	return func(e *Engine[K, U, C]) {
		e.claims = claims
	}
}
func NewEngine[K IKey, U IUser, C IClaims](fns ...EFn[K, U, C]) *Engine[K, U, C] {
	engine := &Engine[K, U, C]{}
	for _, fn := range fns {
		fn(engine)
	}
	return engine
}
func (e *Engine[K, U, C]) Key() K {
	return e.key
}
func (e *Engine[K, U, C]) User() U {
	return e.user
}
func (e *Engine[K, U, C]) Claims() C {
	return e.claims
}
func (e *Engine[K, U, C]) WithKey(key K) IEngine[K, U, C] {
	e.key = key
	return e
}
func (e *Engine[K, U, C]) WithUser(user U) IEngine[K, U, C] {
	e.user = user
	return e
}
func (e *Engine[K, U, C]) WithClaims(claims C) IEngine[K, U, C] {
	e.claims = claims
	return e
}
func (e *Engine[K, U, C]) Generate() (token string, err error) {
	// 签名器
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: e.Key().Private()}, nil)
	if err != nil {
		return "", err
	}

	// 签名 JWT
	j, err := jwt.Signed(signer).Claims(&UserClaims[U, C]{
		user:   e.User(),
		claims: e.Claims(),
	}).CompactSerialize()
	if err != nil {
		return "", err
	}

	// 加密器
	encryptER, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{
		Algorithm: jose.RSA_OAEP_256,
		Key:       e.Key().Public(),
	}, nil)
	if err != nil {
		return "", fmt.Errorf("创建加密器失败: %w", err)
	}

	// 加密签名后的 JWT
	te, err := encryptER.Encrypt([]byte(j))
	if err != nil {
		return "", fmt.Errorf("加密 JWT 失败: %w", err)
	}
	// 使用 CompactSerialize 方法序列化加密对象
	t, err := te.CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("序列化加密 JWT 失败: %w", err)
	}
	return t, err
}
func (e *Engine[K, U, C]) VerifierToken(token string) (UserClaims[U, C], error) {
	// 解密
	object, err := jose.ParseEncrypted(token)
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("解析加密 JWT 失败: %w", err)
	}

	decryptedJWT, err := object.Decrypt(e.Key().Private())
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("解密 JWT 失败: %w", err)
	}

	// 验证签名
	parsedJWT, err := jwt.ParseSigned(string(decryptedJWT))
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("解析签名 JWT 失败: %w", err)
	}

	var parse UserClaims[U, C]
	if err = parsedJWT.Claims(e.Key().Public(), &parse); err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("验证 JWT 签名失败: %w", err)
	}

	// 验证声明
	err = parse.Claims().Validated(jwt.Expected{
		Issuer: e.Claims().GetIssuer(),
		Time:   time.Now().UTC(),
	})
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("JWT 声明无效: %w", err)
	}

	return parse, nil
}
func (e *Engine[K, U, C]) VerifierRefresh(token string) (UserClaims[U, C], error) {
	// 解密
	object, err := jose.ParseEncrypted(token)
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("解析加密 JWT 失败: %w", err)
	}

	decryptedJWT, err := object.Decrypt(e.Key().Private())
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("解密 JWT 失败: %w", err)
	}

	// 验证签名
	parsedJWT, err := jwt.ParseSigned(string(decryptedJWT))
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("解析签名 JWT 失败: %w", err)
	}

	var parse UserClaims[U, C]
	if err := parsedJWT.Claims(e.Key().Public(), &parse); err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("验证 JWT 签名失败: %w", err)
	}

	if !parse.User().IsRefresh() {
		return UserClaims[U, C]{}, fmt.Errorf("验证 JWT 签名失败: 不是刷新类型的 token")
	}

	// 验证声明
	err = parse.Claims().Validated(jwt.Expected{
		Issuer: e.Claims().GetIssuer(),
		Time:   time.Now().UTC(),
	})
	if err != nil {
		return UserClaims[U, C]{}, fmt.Errorf("JWT 声明无效: %w", err)
	}

	return parse, nil
}
