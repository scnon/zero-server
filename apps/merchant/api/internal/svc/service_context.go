package svc

import (
	"context"

	"github.com/scnon/zero-server/apps/merchant/api/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ServiceContext struct {
	context.Context
	Config config.Config

	*gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := gorm.Open(sqlite.Open(c.DataFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,

		DB: conn,
	}
}
