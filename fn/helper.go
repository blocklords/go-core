package fn

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

func Now() int64 {
	return time.Now().UTC().Unix()
}

func NowMs() int64 {
	return time.Now().UTC().UnixMilli()
}

func NowNs() int64 {
	return time.Now().UTC().UnixNano()
}

func Date() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

func HashFromUUID() common.Hash {
	return common.HexToHash(hex.EncodeToString(Sha1([]byte(uuid.New().String()))))
}

type VRS struct {
	r common.Hash `json:"r"`
	s common.Hash `json:"s"`
	v uint8       `json:"v"`
}

func (rsv VRS) R() common.Hash {
	return rsv.r
}

func (rsv VRS) S() common.Hash {
	return rsv.s
}

func (rsv VRS) V() uint8 {
	return rsv.v
}

func (rsv VRS) JSON() map[string]interface{} {
	return map[string]interface{}{
		"v": rsv.v,
		"r": rsv.r,
		"s": rsv.s,
	}
}

func NewVRS(sign []byte) VRS {
	hash := common.Bytes2Hex(sign)
	rS := hash[0:64]
	sS := hash[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vS := hash[128:130]
	vI, _ := strconv.Atoi(vS)

	return VRS{
		r: common.BytesToHash(R[:]),
		s: common.BytesToHash(S[:]),
		v: uint8(vI + 27),
	}
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func RandomCode(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
