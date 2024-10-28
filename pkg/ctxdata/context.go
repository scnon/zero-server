package ctxdata

import (
	"context"

	"gorm.io/gorm"
)

type contextKey string

const DBContextKey = contextKey("db")

// 将 GORM DB 注入 context
func WithDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, DBContextKey, db)
}

// 从 context 中获取 GORM DB 实例
func GetDB(ctx context.Context) *gorm.DB {
	if db, ok := ctx.Value(DBContextKey).(*gorm.DB); ok {
		return db
	}
	return nil
}
