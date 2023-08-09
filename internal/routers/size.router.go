package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	middleware "sasmeka/coffeeshop/internal/middlewares"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func sizes(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/size")

	// dependcy injection
	repo := repositories.New_Sizes(d)
	handler := handlers.New_Sizes(repo)

	route.GET("/", middleware.AuthJwt("admin"), handler.Get_Data_Sizes)
	route.POST("/", middleware.AuthJwt("admin"), handler.Post_Data_Size)
	route.PUT("/:id", middleware.AuthJwt("admin"), handler.Put_Data_Size)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.Delete_Data_Size)

}
