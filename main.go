package main

import (
	"gsn-wordle/internal/dao"
	"gsn-wordle/internal/middleware"
	"gsn-wordle/internal/pkg/config"
	"gsn-wordle/internal/pkg/util"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialization
	util.EnsureNessesaryDirs()
	config.Init()
	dao.InitDB()
	dao.InitRedis()
	authMiddleware := middleware.Auth

	// Create gin-engine and base router-group
	server := gin.New()

	server.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "*")
	})

	r := server.Group("/api")
	r.Use(gin.LoggerWithWriter(os.Stdout, "/api/ping", "/api/ws/p2p")).
		Use(gin.Recovery()).
		Use(authMiddleware.Authenticate)

	//* --------------------------- API Registration --------------------------- *//
	// PING API
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Pong",
			"data": nil,
		})
	})

	// Register API Routes

	// Websocket P2P
	// r.GET("/ws/p2p", rtc.Signal.SignalP2P)

	log.Println("启动成功")
	server.Run(":" + config.Config.Server.Port)
}
