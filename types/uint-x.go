package types

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type UintX struct {
	*big.Int
}

func (x *UintX) MarshalJSON() ([]byte, error) {
	if x == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(x)
}

func (x *UintX) UnmarshalJSON(b []byte) error {
	bg, ok := new(big.Int).SetString(string(b), 10)
	if !ok || bg.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("invalid number: %s", string(b))
	}
	x.Int = bg
	return nil
}
