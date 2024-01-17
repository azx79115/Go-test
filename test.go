package main

import (
	"errors"
	"github.com/azx79115/Go-test.git/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexData struct {
	Title   string
	Content string
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	if in, ok := c.GetPostForm("username"); ok && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}

	if in, ok := c.GetPostForm("password"); ok && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}

	if err := auth.Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}

}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一隻 Gin 專案"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	server.Static("/assets", "./template/assets")
	server.POST("/login", LoginAuth)
	server.GET("/login", LoginPage)
	server.Run(":9090")
}
