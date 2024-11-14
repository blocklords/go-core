package key

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Key struct {
	privateHex string
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    common.Address
}

func LoadKey(privateKey string) *Key {
	private, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}

	public, ok := private.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return &Key{
		privateHex: privateKey,
		privateKey: private,
		publicKey:  public,
		address:    crypto.PubkeyToAddress(*public),
	}
}

func (k *Key) Private() *ecdsa.PrivateKey {
	return k.privateKey
}

func (k *Key) Address() common.Address {
	return k.address
}

func (k *Key) MakeTransactor(gas, chainId *big.Int) *bind.TransactOpts {
	signer := types.LatestSignerForChainID(chainId)
	return &bind.TransactOpts{
		From: k.address,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != k.address {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), k.privateKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
		GasPrice: gas,
		Context:  context.TODO(),
	}
}
