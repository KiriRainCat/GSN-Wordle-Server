package router

import (
	"gsn-wordle/internal/controller"

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
	g.DELETE("/word/:id", c.Delete)
}
