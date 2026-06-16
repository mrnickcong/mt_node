package system

import (
	"mt_node/internal/svc"
	"mt_node/internal/types"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取系统响应
func PingHandler1(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingPangRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		resp := &types.PingPangResponse{
			Message:   req.Pang,
			Timestamp: time.Now().Unix(),
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

// 获取系统时间
func TimeHandler1(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		_, offset := now.Zone()
		resp := &types.SystemTimeResponse{
			Message:        "success",
			Timestamp:      now.Format("2006-01-02 15:04:05"),
			TimeUnix:       now.Unix(),
			TimeUnixNano:   now.UnixNano(),
			Timezone:       now.Location().String(),
			TimezoneOffset: int64(offset),
			DaylightSaving: 0,
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
