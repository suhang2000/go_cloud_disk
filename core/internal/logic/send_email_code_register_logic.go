package logic

import (
	"context"
	"errors"
	"go_cloud_disk/core/define"
	"go_cloud_disk/core/helper"
	"go_cloud_disk/core/models"

	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCodeRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailCodeRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeRegisterLogic {
	return &SendEmailCodeRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailCodeRegisterLogic) SendEmailCodeRegister(req *types.SendEmailCodeRequest) (resp *types.SendEmailResponse, err error) {
	// check if email exists
	cnt, err := models.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("this email is already registered")
		return
	}

	code := helper.RandCode()
	// store email code in Redis
	models.RDB.Set(l.ctx, req.Email, code, define.CodeExpireTime)
	err = helper.SendMailCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
