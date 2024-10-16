package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"skyExplorerBack/src/controller"
	"strings"
	"time"
)

var app *gin.Engine

func Routers() *gin.Engine {

	v1 := app.Group("/api/v1")
	{
		v1.GET("ping", controller.TestHandler)

		v1.GET("processes", controller.ProcessesGetAllHandler)
		v1.GET("processes/:id", controller.ProcessesGetByIdHandler)
		v1.POST("processes/:id/execution", controller.ProcessesExecutionHandler)

		v1.GET("jobs", controller.JobsGetAllHandler)
		v1.GET("jobs/:id", controller.JobsGetByIdHandler)
		v1.DELETE("jobs/:id", controller.JobsCancelHandler)
		v1.GET("jobs/:id/results", controller.JobsCancelHandler)

	}
	return app
}

func init() {
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = io.Discard
	app = gin.Default()

	// 中间件，解决开发时的跨域问题
	app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// 前端界面
	app.Use(static.Serve("/", static.LocalFile("dist", true)))

	// 静态文件的文件夹相对目录
	app.StaticFS("/dist", http.Dir("./dist"))
	// 单文件路径映射
	app.StaticFile("/favicon.ico", "./favicon.ico")
	app.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("dist1/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				_, err := c.Writer.WriteString("Not Found")
				if err != nil {
					logrus.Warning(err)
					return
				}
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			_, err = c.Writer.Write(content)
			if err != nil {
				return
			}
			c.Writer.Flush()
		}
	})
}
