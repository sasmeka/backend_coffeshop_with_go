package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/")

	// dependcy injection
	repo := repositories.New_Auth(d)
	handler := handlers.New_Auth(repo)

	route.POST("/login", handler.Login)
	route.POST("/register", handler.Register)
}
