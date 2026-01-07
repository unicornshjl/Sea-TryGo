// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"sea-try-go/service/user/api/internal/model"
	"sea-try-go/service/user/api/internal/svc"
	"sea-try-go/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteuserLogic {
	return &DeleteuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteuserLogic) Deleteuser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	id := req.Id
	user := model.User{
		Id: id,
	}
	err = l.svcCtx.DB.Delete(&user).Error
	if err != nil {
		return &types.DeleteUserResp{
			Success: false,
		}, err
	}
	return &types.DeleteUserResp{
		Success: true,
	}, nil
}
