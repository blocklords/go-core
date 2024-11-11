package vrs

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

// go test -v .

func TestNewJWT(t *testing.T) {
	decode, err := hexutil.Decode(`0x82a84e9d694fdf77e2cbaf03c4f61d36c339a2bca7e55c62b3b8a03409afa1390f00ee0d700e181c7b93ba4d5d147a1f92dd5e5024602a34d12a2d6823632fd400`)
	if err != nil {
		t.Fatalf("decode error: %+v", err)
	}
	vrs := NewVRS(decode)

	t.Logf("%+v", vrs)
}
