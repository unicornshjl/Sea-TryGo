// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package admin

import (
	"context"

	"sea-try-go/service/admin/api/internal/model"
	"sea-try-go/service/admin/api/internal/svc"
	"sea-try-go/service/admin/api/internal/types"
	"sea-try-go/service/common/cryptx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateAdminReq) (resp *types.CreateAdminResp, err error) {

	password, e := cryptx.PasswordEncrypt(req.Password)
	if e != nil {
		return nil, e
	}
	admin := model.Admin{
		Username:  req.Username,
		Password:  password,
		Email:     req.Email,
		ExtraInfo: req.Extrainfo,
	}
	err = l.svcCtx.DB.Save(&admin).Error
	if err != nil {
		return nil, err
	}
	return &types.CreateAdminResp{
		Id: admin.Id,
	}, nil
}
