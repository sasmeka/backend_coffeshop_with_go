package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	// dependcy injection
	repo := repositories.New_Products(d)
	handler := handlers.New_Products(repo)

	route.GET("/", handler.Get_Data_Products)
	route.POST("/", handler.Post_Data_Product)
	route.PUT("/:id", handler.Put_Data_Product)
	route.DELETE("/:id", handler.Delete_Data_Product)

}
