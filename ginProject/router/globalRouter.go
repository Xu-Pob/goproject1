package router

import (
	"fmt"
	"github.com/Xu-Pob/goproject1/ginProject/controller/api"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWareA(c *gin.Context) {
	fmt.Println("A- init middleware start")
	//中间件中设置值
	c.Set("username", "张三")

	//
	ccp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done in path " + ccp.Request.URL.Path)
	}()
	c.Next()

	fmt.Println("A- init middleware end")
}
func InitMiddleWareB(c *gin.Context) {
	fmt.Println("B- init middleware start")
	//控制器或者中间件中获取值
	username, _ := c.Get("username")
	fmt.Println(username)
	//username 是一个空接口类型，通过转为字符串v, ok := username.(string)
	//v, _ := username.(string)
	//fmt.Println(v)
	c.Next()
	fmt.Println("B- init middleware end")
}

func GlobalRouter(r *gin.Engine) {
	apiRouter := r.Group("/global")
	apiRouter.Use(InitMiddleWareA, InitMiddleWareB)
	{
		// 中间件要放在最后一个回调函数的前面
		apiRouter.GET("/", func(ctx *gin.Context) {
			fmt.Println("首页")
			//设置cookie
			ctx.SetCookie("admin-names", "adminname", 3600, "/global/", "localhost", false, false)

			ctx.String(200, "/global")
		})
		apiRouter.GET("admins", api.AdminController{}.Index)
		apiRouter.GET("getcookie", func(ctx *gin.Context) {
			adname, _ := ctx.Cookie("admin-names")
			ctx.String(200, adname)
		})
	}

}
