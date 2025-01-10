package handler

import (
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
	//a, b := io.ReadAll(c.Request.Body)
	//log.Println("a", a, b)
	//log.Println("a", string(a), b)
	var req BroadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if len(req.SignedTransaction) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SignedTransaction is empty"})
		return
	}
	var signed *types.Transaction
	if util.ReturnErrResp(c, signed.UnmarshalBinary([]byte(req.SignedTransaction))) {
		return
	}
	log.Println("send tx", "signed", signed)
	if err := eth.Cli().SendTransaction(c, signed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to broadcast transaction", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactionHash": signed.Hash().Hex()})
}
