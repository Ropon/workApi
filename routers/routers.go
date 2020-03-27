package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workApi/controller"
)

func SetupRouter() *gin.Engine {
	//记录日志
	//gin.DisableConsoleColor()
	//f, _ := os.Create("./workApi.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//生产模式
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//跨域 频率控制
	r.Use(cors(), limitRate())
	v1 := r.Group("api/v1")
	{
		v1.GET("demo1", controller.Demo1)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"errcode": 4404, "errmsg": "Page not found"})
	})
	return r
}