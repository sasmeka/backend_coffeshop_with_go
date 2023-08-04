package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func users(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	// dependcy injection
	repo := repositories.New_Users(d)
	handler := handlers.New_Users(repo)

	route.GET("/", handler.Get_Data_Users)
	route.POST("/", handler.Post_Data_User)
	route.PUT("/:id", handler.Put_Data_User)
	route.DELETE("/:id", handler.Delete_Data_User)

}
