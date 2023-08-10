package routers

import (
	"sasmeka/coffeeshop/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Routers(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(cors.New(config.CorsConfig))
	// router.Use(middleware.CORSMiddleware)

	auth(router, db)
	users(router, db)
	sizes(router, db)
	delivery_methods(router, db)
	products(router, db)

	return router
}
