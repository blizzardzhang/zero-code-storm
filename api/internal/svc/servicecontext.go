package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero-code-storm/api/internal/config"
	"zero-code-storm/api/internal/middleware"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
	}
}
