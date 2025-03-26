package server

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/yangsai7/demo_api/api"
	"github.com/yangsai7/demo_api/middleware"
	"github.com/yangsai7/demo_api/service"

	_ "net/http/pprof"
)

func NewHTTPServer() *gin.Engine {
	e := gin.Default()
	e.Use(middleware.SetUser)
	pprof.Register(e) // 性能

	api.RegisterUserHTTPServer(e, service.NewUserService())

	return e
}
