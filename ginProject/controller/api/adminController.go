package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	BaseController
}

func (admin AdminController) Index(c *gin.Context) {
	fmt.Println("2-中间件")
	//继承父类的方法
	admin.success(c)
}
func (admin AdminController) Add(c *gin.Context) {
	c.String(200, "api Add")
}
func (admin AdminController) List(c *gin.Context) {
	c.String(200, "api List")
}

func (admin AdminController) Update(c *gin.Context) {
	c.String(200, "api update")
}

func (admin AdminController) Delete(c *gin.Context) {
	c.String(200, "api Delete")
}

//func AdminIndex(c *gin.Context) {
//	c.String(200, "api AdminIndex")
//}

//func AdminAdd(c *gin.Context) {
//	c.String(200, "api AdminAdd")
//}
//
//func AdminList(c *gin.Context) {
//	c.String(200, "api AdminList")
//}
//
//func AdminDelete(c *gin.Context) {
//	c.String(200, "api AdminDelete")
//}
//func AdminUpdate(c *gin.Context) {
//	c.String(200, "api AdminUpdate")
//}
