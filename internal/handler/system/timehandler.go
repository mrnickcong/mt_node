// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mt_node/internal/logic/system"
	"mt_node/internal/svc"
)

// 获取系统时间
func TimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := system.NewTimeLogic(r.Context(), svcCtx)
		resp, err := l.Time()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
