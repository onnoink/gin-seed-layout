//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/onnoink/gin-seed-layout/internal/biz"
	"github.com/onnoink/gin-seed-layout/internal/data"
	"github.com/onnoink/gin-seed-layout/internal/server"
	"github.com/onnoink/gin-seed-layout/internal/service"
)

func wireGin() *gin.Engine {
	panic(wire.Build(
		service.ProviderSet,
		biz.ProviderSet,
		data.NewHelloRepo,
		server.NewServer,
	))
}
