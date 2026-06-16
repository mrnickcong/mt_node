// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mt_node/internal/logic/system"
	"mt_node/internal/svc"
	"mt_node/internal/types"
)

// 获取系统响应
func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingPangRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := system.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
