package client

import (
	"context"

	"zero-code-storm/api/internal/svc"
	"zero-code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClientByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientByIdLogic {
	return &GetClientByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClientByIdLogic) GetClientById(req *types.DetailReq) (resp *types.DetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
