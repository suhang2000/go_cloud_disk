package logic

import (
	"context"
	"errors"
	"go_cloud_disk/core/helper"
	"go_cloud_disk/core/models"

	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1. query user from mysql
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	//fmt.Println("password: ", req.Password)
	//fmt.Println("password md5: ", helper.Md5(req.Password))
	//fmt.Println("user password md5: ", user.Password)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("username or password error")
	}
	// 2. generate token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	// response
	resp = new(types.LoginResponse)
	resp.Token = token
	return
}
