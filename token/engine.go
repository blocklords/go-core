package token

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

type (
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

/**************************************
      Options 模式创建,方便后期扩展
***************************************/

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

/**************************************
      Builder 模式，方便修改 参数值
***************************************/

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

/**************************************
      创建 jwt token 以及验证
***************************************/

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
func (e *Engine[K, U, C]) VerifierToken(token string) (*U, *C, error) {
	// 解密
	object, err := jose.ParseEncrypted(token)
	if err != nil {
		return nil, nil, fmt.Errorf("解析加密 JWT 失败: %w", err)
	}

	decryptedJWT, err := object.Decrypt(e.Key().Private())
	if err != nil {
		return nil, nil, fmt.Errorf("解密 JWT 失败: %w", err)
	}

	// 验证签名
	parsedJWT, err := jwt.ParseSigned(string(decryptedJWT))
	if err != nil {
		return nil, nil, fmt.Errorf("解析签名 JWT 失败: %w", err)
	}

	parse := &UserClaims[U, C]{}
	if err = parsedJWT.Claims(e.Key().Public(), parse); err != nil {
		return nil, nil, fmt.Errorf("验证 JWT 签名失败: %w", err)
	}

	// 验证声明
	err = parse.Claims().Validated(jwt.Expected{
		Issuer: e.Claims().GetIssuer(),
		Time:   time.Now().UTC(),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("JWT 声明无效: %w", err)
	}

	return &parse.user, &parse.claims, nil
}
func (e *Engine[K, U, C]) VerifierRefresh(token string) (*U, *C, error) {
	// 解密
	object, err := jose.ParseEncrypted(token)
	if err != nil {
		return nil, nil, fmt.Errorf("解析加密 JWT 失败: %w", err)
	}

	decryptedJWT, err := object.Decrypt(e.Key().Private())
	if err != nil {
		return nil, nil, fmt.Errorf("解密 JWT 失败: %w", err)
	}

	// 验证签名
	parsedJWT, err := jwt.ParseSigned(string(decryptedJWT))
	if err != nil {
		return nil, nil, fmt.Errorf("解析签名 JWT 失败: %w", err)
	}

	parse := &UserClaims[U, C]{}
	if err = parsedJWT.Claims(e.Key().Public(), parse); err != nil {
		return nil, nil, fmt.Errorf("验证 JWT 签名失败: %w", err)
	}

	if !parse.User().IsRefresh() {
		return nil, nil, fmt.Errorf("验证 JWT 签名失败: 不是刷新类型的 token")
	}

	// 验证声明
	err = parse.Claims().Validated(jwt.Expected{
		Issuer: e.Claims().GetIssuer(),
		Time:   time.Now().UTC(),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("JWT 声明无效: %w", err)
	}

	return &parse.user, &parse.claims, nil
}
