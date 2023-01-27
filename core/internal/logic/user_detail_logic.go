package logic

import (
	"context"
	"errors"
	"go_cloud_disk/core/models"

	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	resp = &types.UserDetailResponse{}
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	resp.Name = user.Name
	resp.Email = user.Email
	return
}
