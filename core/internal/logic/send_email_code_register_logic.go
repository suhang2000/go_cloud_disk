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
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("this email is already registered")
		return
	}
	// 1. generate random code
	code := helper.RandCode()
	// 2. store email code in Redis
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, define.CodeExpireTime)
	// 3. send email
	err = helper.SendMailCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	resp = new(types.SendEmailResponse)
	resp.Message = "the code has been sent to your email, please check your email"
	return
}
