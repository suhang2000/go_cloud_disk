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

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveResponse, err error) {
	// check if resource exists
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return
	}
	if !has {
		return nil, errors.New("resource not found")
	}
	// save into user repository
	ur := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return
	}

	resp = &types.ShareBasicSaveResponse{UserRepositoryIdentity: ur.Identity}
	return
}
