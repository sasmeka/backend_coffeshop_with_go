package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	middleware "sasmeka/coffeeshop/internal/middlewares"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func users(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	// dependcy injection
	repo := repositories.New_Users(d)
	handler := handlers.New_Users(repo)

	route.GET("/", middleware.AuthJwt("admin"), handler.Get_Data_Users)
	route.POST("/", middleware.AuthJwt("admin"), handler.Post_Data_User)
	route.PUT("/:id", middleware.AuthJwt("admin"), handler.Put_Data_User)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.Delete_Data_User)

}
