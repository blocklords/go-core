package token

import (
	"github.com/blocklords/go-core/entity"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

// User 信息相关, IUser
type (
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

// ***************************************
// Option 模式创建 结构体
// ***************************************

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

// ***************************************
// Builder 模式
// ***************************************

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
