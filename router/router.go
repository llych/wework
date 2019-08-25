package router

import (
	"time"
	v1 "wework/controller/v1"
	"wework/extend/conf"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	var (
		r *gin.Engine
	)
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  conf.CORSConf.AllowAllOrigins,
		AllowMethods:     conf.CORSConf.AllowMethods,
		AllowHeaders:     conf.CORSConf.AllowHeaders,
		ExposeHeaders:    conf.CORSConf.ExposeHeaders,
		AllowCredentials: conf.CORSConf.AllowCredentials,
		MaxAge:           conf.CORSConf.MaxAge * time.Hour,
	}))

	apiV1 := r.Group("api/v1")
	//apiV1.Use(middleware.Parms())
	{
		apiV1.GET("/ping", v1.Ping)
		apiV1.POST("/wework", v1.Wework)
	}
	return r
}
