package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认路由
	router := gin.Default()
	// 绑定路由规则和路由函数，访问/index的路由，将由对应的函数去处理
	//router.GET("/index", func(c *gin.Context) {
	//	// 状态码，
	//	c.String(http.StatusOK, "hello world")
	//})
	// 响应json
	router.GET("/json", func(c *gin.Context) {
		// json响应结构体
		type UserInfo struct {
			Username string `json:"username"`
			Age      int    `json:"age"`
			Password string `json:"-"` // 不要转换为字符串
		}
		userInfo := UserInfo{Username: "user", Age: 18, Password: "123456"}
		c.JSON(200, userInfo)
		// json响应map
		var p1Info = map[string]interface{}{
			"p1Name":      "p1",
			"p1Age":       12,
			"p1Passworld": "123456",
		}
		c.JSON(200, p1Info)
		// 直接响应json
		c.JSON(200, gin.H{"json_name": "json", "age": 18, "password": 123456})

	})
	// 响应xml
	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"xml_name": "xml", "age": 18, "password": "123456"})
	})
	// 响应yaml
	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"yaml_name": "yaml", "age": 18, "password": "123456", "data": gin.H{"name": "yaml"}})
	})

	// 加载模版目录下的所有目录文件
	router.LoadHTMLGlob("templates/*")
	// 加载静态文件目录下所有静态文件
	// 在go中没有相对文件的路径，只有相对项目的路径
	// 网页请求静态目录的前缀，第二个参数是一个目录 请求路由-文件目录
	router.StaticFS("/static", http.Dir("static/static"))
	// 网页请求单个文件，请求路由-文件路径
	router.StaticFile("/screen", "./static/screen.jpg")
	// 响应html
	router.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 重定向
	/*
		重定向301和302的区别
		301:表示资源已被永久移动到新的位置，客户端和搜索引擎应使用新的 URL 来访问资源。
		302:表示资源暂时被移动到另一个位置，客户端应继续使用旧的 URL 访问资源。
	*/
	router.GET("/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	/*-------------分割线--------------*/
	// 请求
	// 查询请求
	router.GET("/query", func(c *gin.Context) {
		// 返回查询参数
		user1 := c.Query("user")
		// 返回查询参数和查询结果
		user2, ok := c.GetQuery("user")
		// 返回同一参数的多个查询结果
		user3, ok := c.GetQueryArray("user")
		// 略：user4, ok := c.GetQueryMap("user")
		if !ok {
			fmt.Println("para is null ")
		} else {
			fmt.Println("user is: ", user1, user2, user3)
		}

	})

	// 动态参数
	router.GET("/param/:user_id/:book_id", func(c *gin.Context) {
		userid := c.Param("user_id")
		bookid := c.Param("book_id")
		fmt.Printf("user id is %v ,book id is %v\n", userid, bookid)
	})

	// 表单参数
	router.POST("/form", func(c *gin.Context) {
		fmt.Println("name is ", c.PostForm("name"))
		fmt.Println("names are ", c.PostFormArray("name"))
		fmt.Println("default form is ", c.DefaultPostForm("addr", "0.0.0.0"))
		// 接受所有的form，包括文件
		multi, err := c.MultipartForm()
		fmt.Println("multiform is ", multi, err)
	})
	// 启动监听，gin会把web服务启动在本机的0.0.0.0:8080端口上
	// 启动方式1
	router.Run(":8080")
	// 启动方式2，用原生http启动,router.Run()本质就是http.ListenAndServe进一步封装
	http.ListenAndServe(":8080", router)

}
