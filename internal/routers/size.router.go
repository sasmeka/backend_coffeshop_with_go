package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func sizes(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/size")

	// dependcy injection
	repo := repositories.New_Sizes(d)
	handler := handlers.New_Sizes(repo)

	route.GET("/", handler.Get_Data_Sizes)
	route.POST("/", handler.Post_Data_Size)
	route.PUT("/:id", handler.Put_Data_Size)
	route.DELETE("/:id", handler.Delete_Data_Size)

}
