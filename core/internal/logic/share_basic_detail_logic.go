package logic

import (
	"context"
	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailResponse, err error) {
	// click number add 1
	_, err = l.svcCtx.Engine.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.ShareBasicIdentity)
	if err != nil {
		return
	}
	// get file info
	resp = new(types.ShareBasicDetailResponse)
	_, err = l.svcCtx.Engine.Table("share_basic").
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "repository_pool.identity = share_basic.repository_identity").
		Join("LEFT", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", req.ShareBasicIdentity).Get(resp)
	if err != nil {
		return nil, err
	}

	return
}
