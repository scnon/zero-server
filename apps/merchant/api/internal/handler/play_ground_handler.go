package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/scnon/zero-server/apps/merchant/api/internal/svc"
)

func playGroundHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// l := logic.NewPlayGroundLogic(r.Context(), svcCtx)
		// err := l.PlayGround()
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.Ok(w)
		// }
		playground.Handler("GraphQL", "/api/v1/query").ServeHTTP(w, r.WithContext(svcCtx))
	}
}
