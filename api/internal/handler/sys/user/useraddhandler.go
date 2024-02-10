package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-code-storm/api/internal/logic/sys/user"
	"zero-code-storm/api/internal/svc"
	"zero-code-storm/api/internal/types"
)

func UserAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserAddLogic(r.Context(), svcCtx)
		resp, err := l.UserAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
