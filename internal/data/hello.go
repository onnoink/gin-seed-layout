package data

import (
	"github.com/onnoink/gin-seed-layout/internal/biz"
)

type helloRepo struct {
}

func NewHelloRepo() biz.HelloRepo {
	return &helloRepo{}
}
