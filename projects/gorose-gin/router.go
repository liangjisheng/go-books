package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BootGin 初始化gin
func BootGin() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK,
			`<meta http-equiv="refresh" content="3;URL=/html"><br><br>
<center>
<h1>欢迎来到半小时快速上手golang web编程之用户的增删改查</h1>
</center>
<br><br><br><br>
<center>
<h2 style="color:red">3s后将展示操作界面</h2>
</center>`)
	})

	router.Use(Cors())

	router.GET("/UserAdd", UserAdd)
	router.GET("/UserList", UserList)
	router.GET("/UserEdit", UserEdit)
	router.GET("/UserDelete", UserDelete)

	// 静态文件服务
	router.Static("/html", "./")

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run("127.0.0.1:8080")
	// router.Run(":3000") for a hard coded port
}
