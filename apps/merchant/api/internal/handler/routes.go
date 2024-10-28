// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"github.com/scnon/zero-server/apps/merchant/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ground",
				Handler: playGroundHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/query",
				Handler: queryHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
