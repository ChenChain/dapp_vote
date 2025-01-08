package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"vote_backend/conf"
	"vote_backend/infra/envfile"
	"vote_backend/infra/eth"
	"vote_backend/sol/vote"
)

/**
部署合约， 并将合约地址写入环境变量
*/

func main() {
	ctx := context.TODO()
	auth := eth.GetDefaultAuth(ctx)
	wrap := func(a *bind.TransactOpts, b bind.ContractBackend, c ...interface{}) (common.Address, *types.Transaction, interface{}, error) {
		x, y, _, m := vote.DeployVote(a, b)
		return x, y, nil, m
	}
	address := eth.Deploy(ctx, auth, wrap)
	if err := envfile.UpdateOrAddEnvVariable(conf.Cfg().Vote.Name, address); err != nil {
		log.Fatal(err)
	}
}
