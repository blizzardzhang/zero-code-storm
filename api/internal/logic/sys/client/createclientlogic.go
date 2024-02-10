package client

import (
	"context"

	"zero-code-storm/api/internal/svc"
	"zero-code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClientLogic {
	return &CreateClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateClientLogic) CreateClient(req *types.AddClientReq) (resp *types.AddClientResp, err error) {
	// todo: add your logic here and delete this line

	return
}
