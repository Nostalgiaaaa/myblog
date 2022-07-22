package routes

import (
	v1 "myblog/api/v1"
	"myblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	gin.SetMode(utils.Conf.AppMode)

	r := gin.Default()

	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
		// 用户模块的路由接口
		apiV1.POST("/user/add", v1.AddUser)
		apiV1.GET("/users", v1.GetUsers)
		apiV1.PUT("/user/:id", v1.EditUser)
		apiV1.DELETE("/user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		apiV1.POST("/category/add", v1.AddCategory)
		apiV1.GET("/categories", v1.GetCategories)
		apiV1.PUT("/category/:id", v1.EditCategory)
		apiV1.DELETE("/category/:id", v1.DeleteCategory)
		// 文章模块的路由接口

	}

	r.Run(utils.Conf.HttpPort)

}
