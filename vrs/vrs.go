package vrs

import (
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/iris-contrib/swagger/v12"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

// swagger:model
type VRS struct {
	r common.Hash `json:"r"`
	s common.Hash `json:"s"`
	v uint8       `json:"v"`
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

func (rsv VRS) R() common.Hash {
	return rsv.r
}

func (rsv VRS) S() common.Hash {
	return rsv.s
}

func (rsv VRS) V() uint8 {
	return rsv.v
}

func (rsv VRS) MarshalJSON() ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(map[string]any{
		`v`: rsv.v,
		`r`: rsv.r,
		`s`: rsv.s,
	})
}
