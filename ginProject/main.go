package main

import "github.com/gin-gonic/gin"

// https://gin-gonic.com/docs/ 官方文档
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})

	})
	r.GET("/news", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"news": context.Query("id"),
		})
	})
	//动态路由   http://localhost:8080/users123   ===>     123
	r.GET("/users:pw", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"users": context.Param("pw"),
		})
	})
	r.Run()

}
