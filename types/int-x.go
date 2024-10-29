package types

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type IntX struct {
	*big.Int
}

func (x *IntX) MarshalJSON() ([]byte, error) {
	if x == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(x)
}

func (x *IntX) UnmarshalJSON(b []byte) error {
	bg, ok := new(big.Int).SetString(string(b), 10)
	if !ok {
		return fmt.Errorf("invalid number: %s", string(b))
	}

	x.Int = bg
	return nil
}
