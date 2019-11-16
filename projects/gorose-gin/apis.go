package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserAdd 用户增加
func UserAdd(c *gin.Context) {
	username := c.Query("username")
	age := c.DefaultQuery("age", "0")
	if username == "" {
		c.JSON(http.StatusOK, FailReturn("用户名不能为空"))
		return
	}
	var data = make(map[string]interface{})
	data["username"] = username
	if age != "0" {
		data["age"] = age
	}
	// 执行入库
	affectedRows, err := DB().Table("users").Data(data).Insert()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	// api接口返回
	c.JSON(http.StatusOK, SuccessReturn(affectedRows))
}

// UserDelete 用户删除
func UserDelete(c *gin.Context) {
	uid := c.Query("uid")
	if uid == "" {
		c.JSON(http.StatusOK, FailReturn("用户id不能为空"))
		return
	}
	affectedRows, err := DB().Table("users").Where("uid", uid).Delete()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(affectedRows))
}

// UserEdit 用户编辑
func UserEdit(c *gin.Context) {
	uid := c.Query("uid")
	username := c.DefaultQuery("username", "")
	age := c.DefaultQuery("age", "0")
	if uid == "" {
		c.JSON(http.StatusOK, FailReturn("用户id不能为空"))
		return
	}
	if username == "" && age == "0" {
		c.JSON(http.StatusOK, FailReturn("未修改"))
		return
	}

	var data = make(map[string]interface{})
	if username != "" {
		data["username"] = username
	}
	if age != "0" {
		data["age"] = age
	}
	// 执行入库操作
	affectedRows, err := DB().Table("users").Where("uid", uid).Data(data).Update()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(affectedRows))
}

// UserList 获取用户列表
func UserList(c *gin.Context) {
	// 默认查询50条数据
	userList, err := DB().Table("users").OrderBy("uid desc").Limit(50).Get()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(userList))
}
