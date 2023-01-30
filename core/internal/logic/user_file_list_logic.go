package logic

import (
	"context"
	"go_cloud_disk/core/define"
	"go_cloud_disk/core/models"

	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	uf := make([]*types.UserFile, 0)
	// pagination
	size := req.Size
	if size <= 0 {
		size = define.PageSize
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * size
	// query MySQL
	err = l.svcCtx.Engine.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.name, user_repository.ext, "+
			"repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "repository_pool.identity = user_repository.repository_identity").
		Limit(size, offset).Find(&uf)
	if err != nil {
		return
	}
	// query file count
	cnt, err := l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Count(new(models.UserRepository))
	if err != nil {
		return
	}
	// response
	resp = new(types.UserFileListResponse)
	resp.List = uf
	resp.Count = cnt

	return
}
