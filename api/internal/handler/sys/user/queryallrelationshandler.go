package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-code-storm/api/internal/logic/sys/user"
	"zero-code-storm/api/internal/svc"
	"zero-code-storm/api/internal/types"
)

func QueryAllRelationsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryAllRelationsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewQueryAllRelationsLogic(r.Context(), svcCtx)
		resp, err := l.QueryAllRelations(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
