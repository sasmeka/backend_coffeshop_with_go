package routers

import (
	"sasmeka/coffeeshop/internal/handlers"
	middleware "sasmeka/coffeeshop/internal/middlewares"
	"sasmeka/coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func delivery_methods(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/deliver_method")

	// dependcy injection
	repo := repositories.New_Delivery_Methods(d)
	handler := handlers.New_Delivery_Methods(repo)

	route.GET("/", middleware.AuthJwt("admin"), handler.Get_Data_Delivery_Methods)
	route.POST("/", middleware.AuthJwt("admin"), handler.Post_Data_Delivery_Method)
	route.PUT("/:id", middleware.AuthJwt("admin"), handler.Put_Data_Delivery_Method)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.Delete_Data_Delivery_Method)

}
