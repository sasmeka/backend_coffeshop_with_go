package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Routers(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	users(router, db)
	sizes(router, db)

	return router
}
