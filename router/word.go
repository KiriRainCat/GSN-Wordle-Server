package router

import (
	"gsn-wordle/internal/controller"
	"gsn-wordle/internal/middleware"

	"github.com/gin-gonic/gin"
)

func WordRoutes(r *gin.RouterGroup) {
	g := r.Group("/word-bank")

	// Deps
	c := controller.Word

	// Routes
	g.GET("/words", c.GetList)
	g.GET("/word/:id", c.GetById)
	g.GET("/random-word", c.GetRandomWord)
	g.GET("/word-of-day", c.GetWordOfTheDay)
	g.POST("/word", c.Create)
	g.PUT("/word/:id", c.Update)

	// Admin Auth Required Routes
	g.DELETE("/word/:id", middleware.Auth.AuthenticateAdmin, c.Delete)
	g.GET("/commits", middleware.Auth.AuthenticateAdmin, c.GetCommits)
	g.POST("/commit/approve/:id", middleware.Auth.AuthenticateAdmin, c.ApproveCommit)
	g.PUT("/word/:id/:active", middleware.Auth.AuthenticateAdmin, c.SetActiveState)
}
