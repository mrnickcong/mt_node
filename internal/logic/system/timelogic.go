// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"context"
	"time"

	"mt_node/internal/svc"
	"mt_node/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取系统时间
func NewTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TimeLogic {
	return &TimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TimeLogic) Time() (resp *types.SystemTimeResponse, err error) {
	now := time.Now()
	_, offset := now.Zone()
	resp = &types.SystemTimeResponse{
		Message:        "success",
		Timestamp:      now.Format("2006-01-02 15:04:05"),
		TimeUnix:       now.Unix(),
		TimeUnixNano:   now.UnixNano(),
		Timezone:       now.Location().String(),
		TimezoneOffset: int64(offset),
		DaylightSaving: 0,
	}
	return resp, nil
}
