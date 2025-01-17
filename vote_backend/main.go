package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"vote_backend/handler"
)

// TODO: DDD Model

func CorsMiddleware() gin.HandlerFunc {
	var CorsCfg = cors.DefaultConfig()
	CorsCfg.AddAllowHeaders("Authorization")
	CorsCfg.AllowCredentials = true
	CorsCfg.AllowOriginFunc = func(origin string) bool {
		return true
	}
	return cors.New(CorsCfg)
}

func main() {

	r := gin.Default()
	r.Use(CorsMiddleware())
	r.GET("/get", handler.GetCandidates)
	r.GET("/add", handler.AddCandidate)
	r.GET("/admin_vote", handler.AdminVoteCandidate)
	r.GET("/vote", handler.FrontVote)

	r.POST("broadcast", handler.BroadcastTransaction)
	r.GET("balance", handler.GetEthTokenBalanceOf)

	r.Run(":8080")
}
