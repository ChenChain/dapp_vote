package handler

import (
	"encoding/base64"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"vote_backend/conf"
	"vote_backend/infra/eth"
	"vote_backend/sol/vote"
	"vote_backend/util"
)

var vo *vote.Vote

func init() {
	cli := eth.Load(conf.Cfg().Vote.Address, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return vote.NewVote(address, backend)
	})
	vo = cli.(*vote.Vote)
}

func GetCandidates(c *gin.Context) {
	arr, err := vo.GetCandidates(nil)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, arr)
}

func AdminVoteCandidate(c *gin.Context) {
	address := c.GetString("address")
	index := c.GetInt64("index")
	auth := eth.GetDefaultAuth(c)
	tx, err := vo.Vote(auth, big.NewInt(index))
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gin.H{
		"address": address,
		"tx":      tx.Hash().Hex(),
	})
}

func AddCandidate(c *gin.Context) {
	name := c.Query("name")
	tx, err := vo.AddCandidate(eth.GetDefaultAuth(c), name)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gin.H{
		"tx":   tx.Hash().Hex(),
		"name": name,
	})
}

func FrontVote(c *gin.Context) {
	type resp struct {
		To       string `json:"to"`
		Data     []byte `json:"data"`
		Nonce    uint64 `json:"nonce"`
		Gas      uint64 `json:"gas"`
		GasLimit uint64 `json:"gas_limit"`
	}
	index := c.GetInt64("index")
	publicKeyStr := c.Query("publickey")

	nonce, gas, err := eth.GetNonceGas(c, publicKeyStr)
	if util.ReturnErrResp(c, err) {
		return
	}
	abi, err := vote.VoteMetaData.GetAbi()
	if util.ReturnErrResp(c, err) {
		return
	}
	bs, err := abi.Pack("vote", big.NewInt(index))
	if util.ReturnErrResp(c, err) {
		return
	}
	res := &resp{
		To:       conf.Cfg().Vote.Address,
		Data:     bs,
		Nonce:    nonce,
		Gas:      gas.Uint64(),
		GasLimit: 300000,
	}

	log.Println(res.Data)
	log.Println(base64.StdEncoding.EncodeToString(res.Data))

	c.JSON(200, res)
	log.Printf("FrontVote, index:%+v, pk:%+v,resp:%v", index, publicKeyStr, res)
}
