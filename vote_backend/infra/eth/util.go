package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"vote_backend/conf"
)

func GetKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := crypto.HexToECDSA(conf.Cfg().Account.PriKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return privateKey, publicKeyECDSA
}

func GetDefaultAuth(ctx context.Context) *bind.TransactOpts {
	pri, pub := GetKeys()
	fromAddress := crypto.PubkeyToAddress(*pub)
	nonce, err := Cli().PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := Cli().SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(pri)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice
	return auth
}
