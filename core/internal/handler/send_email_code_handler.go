package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_cloud_disk/core/internal/logic"
	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"
)

func SendEmailCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendEmailCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendEmailCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendEmailCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
