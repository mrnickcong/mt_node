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

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取系统响应
func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.PingPangRequest) (resp *types.PingPangResponse, err error) {
	resp = &types.PingPangResponse{
		Message:   req.Pang,
		Timestamp: time.Now().Unix(),
	}
	return
}
