package router

import "github.com/gin-gonic/gin"

func ApiRounter(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", func(c *gin.Context) {
			c.String(200, "api")
		})
		apiRouter.GET("articles", func(ctx *gin.Context) {
			ctx.String(200, "/api/articles")
		})
	}

}
