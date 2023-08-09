package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Routers(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	auth(router, db)
	users(router, db)
	sizes(router, db)
	delivery_methods(router, db)
	products(router, db)

	return router
}
