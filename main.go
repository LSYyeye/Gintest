package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"` //改成小写
	Desc    string `json:"desc"`
	Content string
}
type Article1 struct {
	Title   string
	Content string
}

// 把时间戳换算成日期
func UnixToTime(timesmap int) string {
	fmt.Println(timesmap)
	t := time.Unix(int64(timesmap), 0)
	return t.Format("2006-01-02 15:04:05")
}
func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "----" + str2
}

func main() {
	// 1.创建路由
	r := gin.Default()

	//自定义模板函数  要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		//注册模板函数
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	//加载模板：html必须要有的，紧挨着路由
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录  第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(200, "get请求，主要用于从服务器中取出资源")
	})
	r.GET("/newss", func(c *gin.Context) {
		c.String(http.StatusOK, "我是新闻页面才怪 122211")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "post请求，主要用于增加数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "put请求，主要用于编辑数据，更新服务器资源")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "delete请求，主要用于删除数据")
	})
	//JSON: map[string]interface{} = gin.H
	//JSON:obj any:不仅可以是空接口也可以是结构体
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin--2",
		})
	})
	//JSONP：主要用于解决跨域问题
	//http://localhost:8000/jsonp?callback=xxx
	//xxx({"title":"我不是标题--jsonp","desc":"不描述","Content":"不测试内容"});
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "我不是标题--jsonp",
			Desc:    "不描述",
			Content: "不测试内容",
		}
		c.JSONP(200, a)
	})
	//XML
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"success": true,
			"msg":     "你好gin--xml",
		})
	})
	// HTML:需要	r.LoadHTMLGlob("templates/*")
	//r:路由；templates：目录名称
	r.GET("/admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "我是后台数据哈哈哈哈哈",
		})
	})
	r.GET("/admin/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/goods.html", gin.H{
			"title": "我是一个gooooooooood页面",
			"price": 20,
		})
	})

	r.GET("/default/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页1111",
			"score": 89,
			"msg":   "我是msg",
			//切片
			"hobby": []string{"吃饭", "睡觉", "写代码"},
			"newsList": []interface{}{
				&Article1{
					Title:   "新闻标题11",
					Content: "新闻内容11",
				},
				&Article1{
					Title:   "新闻标题22",
					Content: "新闻内容22",
				},
			},
			"testSlice": []string{},
			"new": &Article1{
				Title:   "dfgsdf",
				Content: "asgd",
			},
			"date": 1720580688, //时间戳
		})
	})
	r.GET("/default/news", func(c *gin.Context) {
		aa := &Article1{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  aa,
		})
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

//前面是请求，后面是路由，后面随便写
//http.Statusok代表200的状态码
