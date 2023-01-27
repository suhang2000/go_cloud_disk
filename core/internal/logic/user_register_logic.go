package logic

import (
	"context"
	"errors"
	"go_cloud_disk/core/helper"
	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"
	"go_cloud_disk/core/models"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// 1. retrieve code from redis
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("email error or email code expired")
	}
	if code != req.Code {
		return nil, errors.New("code not matching")
	}
	// 2. check if username exists in MySQL
	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		log.Println("MySQL error")
		return nil, errors.New("register user failed")
	}
	if cnt > 0 {
		return nil, errors.New("username already exists")
	}
	// 3. store into MySQL
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	row, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		log.Println("MySQL error")
		return nil, errors.New("register user failed")
	}
	log.Println("insert user row: ", row)

	resp = new(types.UserRegisterResponse)
	resp.Message = "user successfully registered"
	return
}
