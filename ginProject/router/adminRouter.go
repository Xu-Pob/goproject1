package router

import (
	"fmt"
	"github.com/Xu-Pob/goproject1/ginProject/controller/api"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、 记录日志、耗时统计等
// 自定义中间件方法
func InitMiddleWareOne(c *gin.Context) {
	fmt.Println("1，自定义中间件A")
	start := time.Now().UnixNano()
	// 调用c.Next()请求的剩余处理程序
	// c.Next()的语句后面先不执行，先跳转路由匹配的最后一个回调函数执行后，
	// 才执行c.Next()后面的语句

	//终止后面的回调函数,执行这里的语句，注意：放在next后面会不起作用
	//c.Abort()
	c.Next()

	fmt.Println("3a-程序执行完成，计算时间")
	end := time.Now().UnixNano()
	dur := time.Duration(end - start) //==>start = time.now()   dur :=time.since(start)

	fmt.Println(dur.Seconds())
}
func InitMiddleWareTwo(c *gin.Context) {
	fmt.Println("2，自定义中间件B")
	start := time.Now().UnixNano()
	c.Next()
	fmt.Println("3b-程序执行完成，计算时间")
	end := time.Now().UnixNano()
	dur := time.Duration(end - start) //==>start = time.now()   dur :=time.since(start)

	fmt.Println(dur.Seconds())
}
func AdminRouter(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		//中间件放在最后一个回调函数的前面
		adminRouter.GET("/", InitMiddleWareOne, InitMiddleWareTwo, api.AdminController{}.Index)
		//adminRouter.GET("list", func(ctx *gin.Context) {
		//	ctx.String(http.StatusOK, "admin/list")
		//})
		adminRouter.GET(":id", api.AdminController{}.List)
		adminRouter.POST("/", api.AdminController{}.Add)
		adminRouter.PUT(":id", api.AdminController{}.Update)
		adminRouter.DELETE(":id", api.AdminController{}.Delete)
	}

}

//1，自定义中间件A
//2，自定义中间件B
//2-中间件
//3b-程序执行完成，计算时间
//0.0005036
//3a-程序执行完成，计算时间
//0.0017756
