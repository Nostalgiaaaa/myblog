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

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User

	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &data)
	if !valid {
		fmt.Println("app.BindAndValid failed , err:", err)
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Errors()...))
		return
	}

	e := model.CheckUser(data.UserName)

	//fmt.Println(code)
	if e.Code() != errcode.Success.Code() {
		response.ToErrorResponse(errcode.ErrorUserNameUsed.WithDetails())
		return
	}

	model.CreateUser(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  e.Code(),
		"data":    data,
		"message": e.Msg(),
	})

	return

}

// 查询单个用户

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageNum == 0 {
		pageNum = 1
	}

	data := model.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errcode.Success.Code(),
		"data":    data,
		"message": errcode.Success.Msg(),
	})

}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))

	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &data)
	if !valid {
		fmt.Println("app.BindAndValid failed , err:", err)
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Errors()...))
		return
	}

	e := model.CheckUser(data.UserName)

	if e.Code() != errcode.Success.Code() {
		response.ToErrorResponse(errcode.ErrorUserNameUsed.WithDetails())
		return
	}

	e = model.EditUser(id, &data)

	c.JSON(200, gin.H{
		"status":  e.Code(),
		"message": e.Msg(),
	})

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	e := model.DeleteUser(id)
	c.JSON(200, gin.H{
		"status":  e.Code(),
		"message": e.Msg(),
	})
}
