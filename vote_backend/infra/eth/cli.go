package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"vote_backend/conf"
)

var cli *ethclient.Client

func init() {
	client, err := ethclient.Dial(conf.Cfg().URL)
	if err != nil {
		log.Fatal(err)
	}
	cli = client
}

func Cli() *ethclient.Client {
	return cli
}

// Deploy 部署合约, f 合约部署函数
func Deploy(ctx context.Context, auth *bind.TransactOpts,
	f func(*bind.TransactOpts, bind.ContractBackend, ...interface{}) (common.Address, *types.Transaction, interface{}, error),
) string {
	if auth == nil {
		auth = GetDefaultAuth(ctx)
	}
	add, _, _, err := f(auth, cli)
	if err != nil {
		log.Fatal(err)
	}
	return add.Hex()
}

// Load 从合约地址 加载合约
func Load(address string, loadFunc func(common.Address, bind.ContractBackend) (interface{}, error)) interface{} {
	add := common.HexToAddress(address)
	instance, err := loadFunc(add, Cli())
	if err != nil {
		log.Fatal(err)
	}
	return instance
}
