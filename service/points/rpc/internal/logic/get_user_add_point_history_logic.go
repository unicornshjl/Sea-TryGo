package logic

import (
	"context"

	"sea-try-go/service/points/rpc/internal/svc"
	"sea-try-go/service/points/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddPointHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddPointHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddPointHistoryLogic {
	return &GetUserAddPointHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAddPointHistoryLogic) GetUserAddPointHistory(in *__.GetUserAddPointHistoryReq) (*__.GetUserAddPointHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &__.GetUserAddPointHistoryResp{}, nil
}
