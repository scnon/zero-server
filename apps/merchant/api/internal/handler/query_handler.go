package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/scnon/zero-server/apps/merchant/api/graph"
	"github.com/scnon/zero-server/apps/merchant/api/internal/svc"
	"github.com/scnon/zero-server/apps/merchant/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func queryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		srv := handler.NewDefaultServer(
			graph.NewExecutableSchema(graph.Config{
				// Resolvers: &resolver.Resolver{},
			}),
		)
		srv.ServeHTTP(w, r.WithContext(svcCtx))
	}
}
