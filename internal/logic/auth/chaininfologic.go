// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"mt_node/internal/svc"
	"mt_node/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChainInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 公链信息
func NewChainInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChainInfoLogic {
	return &ChainInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChainInfoLogic) ChainInfo() (resp []*types.ChainInfo, err error) {
	//查询所有的ChainInfo信息
	chainInfos, err := l.svcCtx.ChainInfoModel.FindAll(l.ctx)
	if err != nil {
		return
	}
	return chainInfos, nil
}
