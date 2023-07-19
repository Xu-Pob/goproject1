package main

import (
	"github.com/Xu-Pob/goproject1/ginProject/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// https://gin-gonic.com/docs/ 官方文档
func main() {
	r := gin.Default()
	//模式匹配
	// 获取 main.go 文件的绝对路径
	absPath, _ := filepath.Abs("./main.go")
	// 拼接模板路径
	templatePath := filepath.Join(filepath.Dir(absPath), "ginProject/templates/*/*.html")

	r.LoadHTMLGlob(templatePath)

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
	//c.String() c.JSON() c.JSONP() c.XML() c.HTML()
	r.GET("/news/string", func(context *gin.Context) {
		aid := context.Query("id")
		context.String(http.StatusOK, aid)
	})
	//返回JSON
	r.GET("/structJson", func(c *gin.Context) {
		var msg struct {
			Username string `json:"username"`
			Msg      string `json:"msg"`
			Age      int    `json:"age"`
		}
		msg.Username = "Pobo"
		msg.Msg = "New Lession"
		msg.Age = 14
	})
	//返回JSOPN
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(200, data)
	})
	//渲染模板 todo:这里存在相同文件名找不对位置问题
	r.GET("/page1", func(c *gin.Context) {
		c.HTML(200, "page1/index.html", map[string]interface{}{"title": "前台首页"})
	})

	//GET /user?uid=20&page=1
	r.GET("/users", func(c *gin.Context) {
		uid := c.Query("uid")
		//给了个默认值
		page := c.DefaultQuery("page", "0")
		c.String(200, "uid=%v,page=%v", uid, page)
	})
	//接受表单传的值
	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	r.GET("/userinfo", func(context *gin.Context) {
		var userinfo Userinfo
		//绑定结构体，引用传递&    返回json   {"user":"zhangsan","password":"123456"}
		if err := context.ShouldBind(&userinfo); err == nil {
			context.JSON(http.StatusOK, userinfo)
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	r.POST("/doLogin", func(c *gin.Context) {
		var userinfo Userinfo
		//Post绑定结构体
		if err := c.ShouldBind(&userinfo); err == nil {
			c.JSON(http.StatusOK, userinfo)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	//content-type:application/xml
	r.POST("/xml", func(c *gin.Context) {
		var article Article
		if err := c.ShouldBindXML(&article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	//路由分组 分离
	//apiRouter := r.Group("/api")
	//{
	//	apiRouter.GET("/", func(ctx *gin.Context) {
	//		ctx.String(200, "api")
	//	})
	//
	//	apiRouter.GET("articles", func(ctx *gin.Context) {
	//		ctx.String(200, "/api/articles")
	//	})
	//}
	router.AdminRouter(r)
	router.ApiRounter(r)
	r.Run(":8080")

}

type Userinfo struct {
	Username string `json:"user" form:"username"`
	Password string `json:"password" form:"password"`
}
type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}
