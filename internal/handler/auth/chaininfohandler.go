// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mt_node/internal/logic/auth"
	"mt_node/internal/svc"
)

// 公链信息
func ChainInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewChainInfoLogic(r.Context(), svcCtx)
		resp, err := l.ChainInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
