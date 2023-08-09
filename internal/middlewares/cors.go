package middleware

import "github.com/gin-gonic/gin"

// func CORSMiddleware(ctx *gin.Context) {
// 	ctx.Header("Access-Control-Allow-Origin", "*")
// 	ctx.Header("Access-Control-Allow-Credentials", "true")
// 	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 	ctx.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

// 	if ctx.Request.Method == "OPTIONS" {
// 		ctx.AbortWithStatus(204)
// 		return
// 	}

// 	ctx.Next()
// }

func CORSMiddleware(ctx *gin.Context) {
	// ctx.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
