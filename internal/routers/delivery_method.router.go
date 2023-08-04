package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func delivery_methods(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/deliver_method")

	// dependcy injection
	repo := repositories.New_Delivery_Methods(d)
	handler := handlers.New_Delivery_Methods(repo)

	route.GET("/", handler.Get_Data_Delivery_Methods)
	route.POST("/", handler.Post_Data_Delivery_Method)
	route.PUT("/:id", handler.Put_Data_Delivery_Method)
	route.DELETE("/:id", handler.Delete_Data_Delivery_Method)

}
