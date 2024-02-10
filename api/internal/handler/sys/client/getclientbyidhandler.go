package client

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-code-storm/api/internal/logic/sys/client"
	"zero-code-storm/api/internal/svc"
	"zero-code-storm/api/internal/types"
)

func GetClientByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := client.NewGetClientByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetClientById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
