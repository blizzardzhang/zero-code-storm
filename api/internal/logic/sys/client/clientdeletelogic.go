package client

import (
	"context"

	"zero-code-storm/api/internal/svc"
	"zero-code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientDeleteLogic {
	return &ClientDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientDeleteLogic) ClientDelete(req *types.DeleteClientrReq) (resp *types.DeleteClientResp, err error) {
	// todo: add your logic here and delete this line

	return
}
