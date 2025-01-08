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
	r.GET("/vote", handler.VoteCandidate)

	r.Run(":8080")
}
