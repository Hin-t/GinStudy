# GinStudy

### 1. 
1. 创建一个默认路由

```go
// 创建一个默认路由
router := gin.Default()
```
2. 绑定路由规则和路由函数，访问/index的路由，将由对应的函数去处理
```go
router.GET("/index", func(c *gin.Context) {
	// 状态码，用于返回状态
	c.String(http.StatusOK, "hello world")
})
```
3. 启动监听
启动监听，gin会把web服务启动在本机的0.0.0.0:8080端口上
```go
// 启动方式1
router.Run(":8080")
// 启动方式2，用原生http启动,router.Run()本质就是http.ListenAndServe进一步封装
http.ListenAndServe(":8080", router)

```
4. 响应json、xml、yaml数据
 ```go
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
c.YAML(http.StatusOK, gin.H{"yaml_name": "yaml", "age": 18, "password": "123456"})
})
```
5. 加载html文件和静态文件
```go
// 加载模版目录下的所有目录文件
router.LoadHTMLGlob("templates/*")
// 加载静态文件目录下所有静态文件
// 在go中没有相对文件的路径，只有相对项目的路径
router.StaticFile("/static/screen.jpg", "./static/screen.jpg")
// 响应html
router.GET("/html", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
})
```
6. 