package handler

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"vote_backend/conf"
	"vote_backend/infra/eth"
	"vote_backend/sol/vote"
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

func VoteCandidate(c *gin.Context) {
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
