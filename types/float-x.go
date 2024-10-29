package types

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type FloatX decimal.Decimal

func (x *FloatX) MarshalJSON() ([]byte, error) {
	if x == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(x)
}

func (x *FloatX) UnmarshalJSON(b []byte) error {
	dc, err := decimal.NewFromString(string(b))
	if err != nil {
		return err
	}

	*x = FloatX(dc)
	return nil
}

func (x *FloatX) Decimal() decimal.Decimal {
	return decimal.Decimal(*x)
}

func (x *FloatX) Float64() float64 {
	return decimal.Decimal(*x).InexactFloat64()
}

func (x *FloatX) Float32() float32 {
	return float32(x.Float64())
}
