// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"mt_node/internal/svc"
	"mt_node/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Mt_node1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMt_node1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Mt_node1Logic {
	return &Mt_node1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Mt_node1Logic) Mt_node1(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
