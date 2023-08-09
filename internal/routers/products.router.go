package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	middleware "sasmeka/coffeeshop/internal/middlewares"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	// dependcy injection
	repo := repositories.New_Products(d)
	handler := handlers.New_Products(repo)

	route.GET("/", middleware.AuthJwt("admin", "user"), handler.Get_Data_Products)
	route.POST("/", middleware.AuthJwt("admin"), handler.Post_Data_Product)
	route.PUT("/:id", middleware.AuthJwt("admin"), handler.Put_Data_Product)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.Delete_Data_Product)

}
