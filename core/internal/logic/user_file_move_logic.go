package logic

import (
	"context"
	"errors"
	"go_cloud_disk/core/models"

	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	// check if parent folder exists
	parentData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("parent folder not found")
	}
	// check if parent_id item is a folder
	if parentData.Ext != "" {
		return nil, errors.New("parent_id item is not a folder")
	}
	// move file
	affected, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(&models.UserRepository{ParentId: int64(parentData.Id)})
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("no column affected, file not found")
	}
	resp = &types.UserFileMoveResponse{Message: "move file successfully"}
	return
}
