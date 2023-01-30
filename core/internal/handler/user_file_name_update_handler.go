package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_cloud_disk/core/internal/logic"
	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"
)

func UserFileNameUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileNameUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFileNameUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserFileNameUpdate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
