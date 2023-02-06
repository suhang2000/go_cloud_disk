package logic

import (
	"context"
	"go_cloud_disk/core/models"

	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteResponse, err error) {
	_, err = l.svcCtx.Engine.Where("identity= ? AND user_identity = ?", req.Identity, userIdentity).Delete(new(models.UserRepository))
	if err != nil {
		return
	}
	resp = &types.UserFileDeleteResponse{Message: "delete successfully"}
	return
}
