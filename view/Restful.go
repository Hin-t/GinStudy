package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func main() {
	router := gin.Default()

	// 文章列表
	router.GET("/articles", func(c *gin.Context) {
		articles := []ArticleModel{
			{Title: "go语言入门", Content: "入门指南"},
			{Title: "信息安全数学基础", Content: "中国剩余定理"},
			{Title: "python语言入门", Content: "这是python语言入门"},
		}
		c.JSON(200, Response{0, articles, "成功"})
	})
	// 文章详情
	router.GET("/articlesdetail/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		articles := ArticleModel{Title: "go语言入门", Content: "入门指南"}
		c.JSON(200, Response{0, articles, "成功"})
	})
	// 创建文章
	router.POST("/createarticles", func(c *gin.Context) {
		var articles ArticleModel
		body, _ := c.GetRawData()
		json.Unmarshal(body, &articles)
		fmt.Println(articles)
	})
	// 更新文章
	router.PUT("/updatearticles/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		var articles ArticleModel
		body, _ := c.GetRawData()
		json.Unmarshal(body, &articles)
		fmt.Println(articles)
		c.JSON(200, Response{0, articles, "修改成功"})
	})
	// 删除文章
	router.DELETE("/deletearticles/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		c.JSON(200, Response{0, nil, "删除成功"})
	})
	router.Run(":8080")
}
