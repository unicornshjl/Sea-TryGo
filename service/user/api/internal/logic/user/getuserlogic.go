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

type GetuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserLogic {
	return &GetuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserLogic) Getuser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	id := req.Id
	user := model.User{}
	err = l.svcCtx.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &types.GetUserResp{
			Found: false,
		}, err
	}
	userInfo := types.UserInfo{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Extrainfo: user.ExtraInfo,
	}
	return &types.GetUserResp{
		User:  userInfo,
		Found: true,
	}, nil
}
