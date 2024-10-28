package logic

import (
	"context"

	"github.com/scnon/zero-server/apps/merchant/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PlayGroundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayGroundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayGroundLogic {
	return &PlayGroundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayGroundLogic) PlayGround() error {
	// todo: add your logic here and delete this line

	return nil
}
