package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRouter(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", func(c *gin.Context) {
			c.String(200, "admin")
		})
		adminRouter.GET("list", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "admin/list")
		})
	}
}
