package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	v1 "github.com/onnoink/gin-seed-layout/api/helloworld/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewServer)

type Service interface {
	Hello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error)
}

func RegisterHttpServer(r *gin.Engine, svc Service) {
	apiV1 := r.Group("/api/v1")
	// 微信code登录
	apiV1.GET("/helloworld", _HelloWorldHandler(svc))
}

func _HelloWorldHandler(svc Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req v1.HelloRequest
		if ctx.ShouldBind(&req) != nil {
			ctx.JSON(400, "Bad Params")
		}
		rv, err := svc.Hello(ctx, &req)
		if err != nil {
			ctx.JSON(400, err.Error())
		}
		ctx.JSON(200, rv)
	}
}

func NewServer(svc Service) *gin.Engine {
	// 创建默认路由引擎
	r := gin.Default()
	RegisterHttpServer(r, svc)
	return r
}
