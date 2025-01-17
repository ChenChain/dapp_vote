package handler

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vote_backend/infra/eth"
	"vote_backend/util"
)

// BroadcastRequest 用于接收签名后的交易
type BroadcastRequest struct {
	SignedTransaction string `json:"signed_transaction"`
}

// BroadcastTransaction 广播交易
func BroadcastTransaction(c *gin.Context) {
	var req BroadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if len(req.SignedTransaction) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SignedTransaction is empty"})
		return
	}

	// 解析签名交易
	rawTx, err := hexutil.Decode(req.SignedTransaction)
	if err != nil {
		log.Println("Invalid signed transaction:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signed transaction"})
		return
	}
	log.Println("BroadcastTransaction get rawTx", rawTx)

	// 发送交易

	signed := &types.Transaction{}
	if util.ReturnErrResp(c, signed.UnmarshalBinary(rawTx)) {
		return
	}
	log.Println("send tx", "signed", signed)
	if err := eth.Cli().SendTransaction(c, signed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to broadcast transaction", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactionHash": signed.Hash().Hex()})
}

func GetEthTokenBalanceOf(c *gin.Context) {
	publicKeyStr := c.Query("publickey")

	balance, err := eth.Cli().BalanceAt(c, common.HexToAddress(publicKeyStr), nil)
	if util.ReturnErrResp(c, err) {
		return
	}
	util.ReturnResp(c, map[string]int64{
		"balance": balance.Int64(),
	})
}
