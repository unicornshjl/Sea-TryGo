// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"errors"

	"sea-try-go/service/user/api/internal/model"
	"sea-try-go/service/user/api/internal/svc"
	"sea-try-go/service/user/api/internal/types"
	"sea-try-go/service/user/common/cryptx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	user := model.User{}

	isExist := l.svcCtx.DB.Where("username = ?", req.Username).First(&user).Error
	if isExist == nil {
		return nil, errors.New("用户名已存在")
	}

	TruePassword, e := cryptx.PassWordEncrypt(req.Password)
	if e != nil {
		return nil, e
	}

	user = model.User{
		Username:  req.Username,
		Password:  TruePassword,
		Email:     req.Email,
		ExtraInfo: req.Extrainfo,
	}
	result := l.svcCtx.DB.Create(&user)

	if result != nil {
		return nil, result.Error
	}

	return &types.CreateUserResp{
		Id: user.Id,
	}, nil
}
