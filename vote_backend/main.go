package main

import (
	"github.com/gin-gonic/gin"
	"vote_backend/handler"
)

// TODO: DDD Model

func main() {

	r := gin.Default()

	r.GET("/get", handler.GetCandidates)
	r.GET("/add", handler.AddCandidate)
	r.GET("/admin_vote", handler.AdminVoteCandidate)
	r.GET("/vote", handler.FrontVote)

	r.POST("broadcast", handler.BroadcastTransaction)

	r.Run(":8080")
}
