// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fmt"
	"net/http"
	"time"

	"mt_node/internal/svc"
	"mt_node/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PingPangHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingPangRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		//获取路径参数  127.0.0.1:8888/ping/pang
		httpx.OkJsonCtx(r.Context(), w, &types.PingPangResponse{
			Message:   req.Pang,
			Timestamp: time.Now().Unix(),
		})
	}
}

func SystemTimeHandler1(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		local := now.Local()

		// 获取时区名称和偏移量
		zoneName, offset := local.Zone()

		httpx.OkJsonCtx(r.Context(), w, &types.SystemTimeResponse{
			Message:        "success",
			Timestamp:      local.Format("2006-01-02 15:04:05"),
			TimeUnix:       local.Unix(),
			TimeUnixNano:   local.UnixNano(),
			Timezone:       zoneName,
			TimezoneOffset: int64(offset),
			DaylightSaving: 0, // Go 标准库不直接提供夏令时信息，设为 0
		})
	}
}

func SystemTimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		local := now.Local()

		// 获取时区名称和偏移量
		zoneName, offset := local.Zone()

		// 生成更明确的时区标识
		timezoneStr := formatTimezone(zoneName, offset)

		httpx.OkJsonCtx(r.Context(), w, &types.SystemTimeResponse{
			Message:        "success",
			Timestamp:      local.Format("2006-01-02 15:04:05"),
			TimeUnix:       local.Unix(),
			TimeUnixNano:   local.UnixNano(),
			Timezone:       timezoneStr,
			TimezoneOffset: int64(offset),
			DaylightSaving: 0,
		})
	}
}

// formatTimezone 格式化时区信息，避免歧义
func formatTimezone(zoneName string, offset int) string {
	hours := offset / 3600
	minutes := (offset % 3600) / 60

	// 如果偏移量是中国标准时间，明确标注
	if offset == 28800 {
		return "CST (UTC+8, China Standard Time)"
	}

	sign := "+"
	if offset < 0 {
		sign = "-"
		offset = -offset
	}

	return zoneName + " (UTC" + sign +
		fmt.Sprintf("%d:%02d", hours, minutes) + ")"
}
