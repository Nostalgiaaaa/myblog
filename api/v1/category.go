package v1

import (
	"fmt"
	"myblog/model"
	"myblog/pkg/app"
	"myblog/utils/errcode"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddCategory  添加分类
func AddCategory(c *gin.Context) {
	var data model.Category

	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &data)
	if !valid {
		fmt.Println("app.BindAndValid failed , err:", err)
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Errors()...))
		return
	}

	e := model.CheckCategory(data.Name)

	//fmt.Println(code)
	if e.Code() != errcode.Success.Code() {
		response.ToErrorResponse(errcode.ErrorCateNameUsed.WithDetails())
		return
	}

	e = model.CreateCategory(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  e.Code(),
		"data":    data,
		"message": e.Msg(),
	})

	return

}

// 查询分类下的所有文章

// GetCategories  查询分类列表
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageNum == 0 {
		pageNum = 1
	}

	data := model.GetCate(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errcode.Success.Code(),
		"data":    data,
		"message": errcode.Success.Msg(),
	})

}

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))

	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &data)
	if !valid {
		fmt.Println("app.BindAndValid failed , err:", err)
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Errors()...))
		return
	}

	e := model.CheckCategory(data.Name)

	if e.Code() != errcode.Success.Code() {
		response.ToErrorResponse(e.WithDetails())
		return
	}

	e = model.EditCategory(id, &data)

	c.JSON(200, gin.H{
		"status":  e.Code(),
		"message": e.Msg(),
	})

}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	e := model.DeleteCategory(id)
	c.JSON(200, gin.H{
		"status":  e.Code(),
		"message": e.Msg(),
	})
}
