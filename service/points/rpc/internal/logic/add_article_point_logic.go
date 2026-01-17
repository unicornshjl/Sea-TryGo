package logic

import (
	"context"

	"sea-try-go/service/points/rpc/internal/svc"
	"sea-try-go/service/points/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticlePointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticlePointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticlePointLogic {
	return &AddArticlePointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddArticlePointLogic) AddArticlePoint(in *__.AddArticlePointReq) (*__.AddArticlePointResp, error) {
	// todo: add your logic here and delete this line

	return &__.AddArticlePointResp{}, nil
}
