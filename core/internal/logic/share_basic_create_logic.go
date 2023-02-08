package logic

import (
	"context"
	"errors"
	"go_cloud_disk/core/helper"
	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"
	"go_cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	// check if file exists
	userRepository := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity=?", req.UserRepositoryIdentity).Get(userRepository)
	if err != nil {
		return
	}
	if !has {
		return nil, errors.New("this file is not in the user repository")
	}
	// create share basic
	shareBasic := &models.ShareBasic{
		Identity:               helper.UUID(),
		UserIdentity:           userIdentity,
		RepositoryIdentity:     userRepository.RepositoryIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	_, err = l.svcCtx.Engine.Insert(shareBasic)
	if err != nil {
		return nil, err
	}

	resp = &types.ShareBasicCreateResponse{Identity: shareBasic.Identity}
	return
}
